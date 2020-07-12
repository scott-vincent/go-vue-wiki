package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/scott-vincent/go-vue-wiki/backend/page"
)

func handleAllPages(w http.ResponseWriter, r *http.Request) {
	titles, err := page.GetTitles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(titles)
}

func handleOnePage(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/pages/"):]
	page, err := page.Load(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(page)
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
		p := &page.Page{Body: []byte(body), Error: "Page must have a title"}
		return p
	}

	// If page title has changed, make sure it is valid
	if newTitle != oldTitle {
		err := page.ValidateNewPage(newTitle)
		if err != nil {
			// Redisplay the edit page and show the error
			p := &page.Page{Title: oldTitle, Body: []byte(body), Error: err.Error()}
			return p
		}

		// Delete the old page
		page.Delete(oldTitle)
	}

	p := &page.Page{Title: newTitle, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return p
	}

	return nil
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/delete/"):]
	page.Delete(title)
}

func main() {
	// Serve static files for frontend
	staticFiles := http.FileServer(http.Dir("../frontend/dist"))
	http.Handle("/", staticFiles)

	// Server endpoints for backend
	http.HandleFunc("/pages", handleAllPages)
	http.HandleFunc("/pages/", handleOnePage)

	// Start the server
	port := 8080
	fmt.Println("Server listening on port", port)
	log.Panic(
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil),
	)
}
