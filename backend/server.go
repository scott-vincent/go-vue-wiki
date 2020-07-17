package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
	if p.Title != oldTitle {
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
	if oldTitle != "*" && oldTitle != p.Title {
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
	router.Methods("OPTIONS").PathPrefix("/wiki-server") // CORS doesn't work without this!

	// Backend - Server endpoints
	router.Methods("GET").PathPrefix("/wiki-server/pages/{title}").HandlerFunc(getPage)
	router.Methods("POST").PathPrefix("/wiki-server/pages/{title}").HandlerFunc(savePage)
	router.Methods("DELETE").PathPrefix("/wiki-server/pages/{title}").HandlerFunc(deletePage)
	router.Methods("GET").PathPrefix("/wiki-server/pages").HandlerFunc(getPages)

	// Frontend - Serve static files
	appDir := getAppDir()
	router.Methods("GET").PathPrefix("/").Handler(http.FileServer(http.Dir(appDir)))

	// Start the server
	port := 8080
	fmt.Println("Server listening on port", port)
	absAppDir, _ := filepath.Abs(appDir)
	fmt.Println("Serving static files from", absAppDir)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
