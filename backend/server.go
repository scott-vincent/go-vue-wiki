package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/scott-vincent/go-vue-wiki/backend/page"
)

///
// GET /pages/:title
///
func getPage(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	title := pathParams["title"]

	p, err := page.Load(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(p)
}

///
// POST /pages/:title
///
func savePage(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	oldTitle := pathParams["title"]

	var p page.Page
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If title changed make sure new page does not already exist
	if (p.Title != oldTitle) {
		err := page.ValidateNewPage(p.Title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	err = p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If title changed delete old page
	if (oldTitle != "*" && oldTitle != p.Title) {
		page.Delete(oldTitle)
	}

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
	router.Methods("POST").PathPrefix("/pages/{title}").HandlerFunc(savePage)
	router.Methods("DELETE").PathPrefix("/pages/{title}").HandlerFunc(deletePage)
	router.Methods("GET").PathPrefix("/pages").HandlerFunc(getPages)

	// Start the server
	port := 8080
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
