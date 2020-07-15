package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/scott-vincent/go-vue-wiki/backend/page"
)

///
// GET /pages
///
func getPages(w http.ResponseWriter, r *http.Request) {
	titles, err := page.GetTitles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(titles)
}

///
// GET /pages/:title
///
func getPage(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	title := pathParams["title"]

	page, err := page.Load(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(page)
}

///
// POST /pages/:title
///
func savePage(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	title := pathParams["title"]

	_ = title

	json.NewEncoder(w).Encode("OK")
}

///
// DELETE /pages/:title
///
func deletePage(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	title := pathParams["title"]

	err := page.Delete(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("OK")
}

func editHandler(w http.ResponseWriter, r *http.Request) *page.Page {
	title := r.URL.Path[len("/edit/"):]
	p, err := page.Load(title)
	if err != nil {
		p = &page.Page{Title: title}
	}
	return p
}

func saveHandler(w http.ResponseWriter, r *http.Request) *page.Page {
	oldTitle := r.URL.Path[len("/save/"):]
	newTitle := strings.TrimSpace(r.FormValue("title"))
	body := r.FormValue("body")

	if newTitle == "" {
		p := &page.Page{Body: body, Error: "Page must have a title"}
		return p
	}

	// If page title has changed, make sure it is valid
	if newTitle != oldTitle {
		err := page.ValidateNewPage(newTitle)
		if err != nil {
			// Redisplay the edit page and show the error
			p := &page.Page{Title: oldTitle, Body: body, Error: err.Error()}
			return p
		}

		// Delete the old page
		page.Delete(oldTitle)
	}

	p := &page.Page{Title: newTitle, Body: body}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return p
	}

	return nil
}

func getAppDir() string {
	prodDir := "dist"
	devDir := "../frontend/dist"

	stat, err := os.Stat(prodDir)
	if err == nil && stat.IsDir() {
		return prodDir
	}

	stat, err = os.Stat(devDir)
	if err == nil && stat.IsDir() {
		return devDir
	}

	log.Fatal("App not found looking for folder ", prodDir, " or ", devDir)
	return ""
}

func main() {
	// Use Gorilla Mux
	router := mux.NewRouter()

	// Allow CORS on localhost
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:8081"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowCredentials(),
	)
	router.Use(cors)
	router.Methods("OPTIONS").PathPrefix("/") // CORS doesn't work without this!

	// Serve static files for frontend
	router.Handle("/", http.FileServer(http.Dir(getAppDir())))

	// Server endpoints for backend
	router.Methods("GET").PathPrefix("/pages/{title}").HandlerFunc(getPage)
	router.Methods("GET").PathPrefix("/pages").HandlerFunc(getPages)
	router.Methods("POST").PathPrefix("/pages/{title}").HandlerFunc(savePage)
	router.Methods("DELETE").PathPrefix("/pages/{title}").HandlerFunc(deletePage)

	// Start the server
	port := 8080
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
