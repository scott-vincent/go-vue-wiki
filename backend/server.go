package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/scott-vincent/go-vue-wiki/backend/page"
)

func allowCorsOnLocalhost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
}

func handleAllPages(w http.ResponseWriter, r *http.Request) {
	allowCorsOnLocalhost(w, r)
	titles, err := page.GetTitles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(titles)
}

func handleOnePage(w http.ResponseWriter, r *http.Request) {
	allowCorsOnLocalhost(w, r)
	title := r.URL.Path[len("/pages/"):]
	page, err := page.Load(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(page)
}

func editHandler(w http.ResponseWriter, r *http.Request) *page.Page {
	allowCorsOnLocalhost(w, r)
	title := r.URL.Path[len("/edit/"):]
	p, err := page.Load(title)
	if err != nil {
		p = &page.Page{Title: title}
	}
	return p
}

func saveHandler(w http.ResponseWriter, r *http.Request) *page.Page {
	allowCorsOnLocalhost(w, r)
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
	allowCorsOnLocalhost(w, r)
	title := r.URL.Path[len("/delete/"):]
	page.Delete(title)
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

	log.Panic("App not found looking for folder ", prodDir, " or ", devDir)
	return ""
}

func main() {
	// Serve static files for frontend
	staticFiles := http.FileServer(http.Dir(getAppDir()))
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
