package page

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const pageFolder = "data"

// Page stores a web page
type Page struct {
	Title string
	Body  string
}

func getFilename(title string) (string, error) {
	if strings.ContainsAny(title, "#?/\\*\"") {
		return "", fmt.Errorf("Page name not allowed - Must not contain special characters")
	}

	return pageFolder + "/" + title + ".txt", nil
}

// Save page
func (p *Page) Save() error {
	filename, err := getFilename(p.Title)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

// Load page
func Load(title string) (*Page, error) {
	filename, err := getFilename(title)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: string(body)}, nil
}

// GetTitles reads all filenames in the page folder
func GetTitles() ([]string, error) {
	files, err := ioutil.ReadDir(pageFolder)
	if err != nil {
		return nil, err
	}

	var titles []string
	for _, file := range files {
		title := file.Name()

		// Remove .txt extension
		titles = append(titles, title[0:len(title)-4])
	}

	sort.Strings(titles)
	return titles, nil
}

// Delete the page with the specified title
func Delete(title string) error {
	filename, err := getFilename(title)
	if err != nil {
		return err
	}

	return os.Remove(filename)
}

// ValidateNewPage returns error if the filename is not valid or already exists
func ValidateNewPage(title string) error {
	filename, err := getFilename(title)
	if err != nil {
		return err
	}

	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		return nil
	} else if err == nil {
		return fmt.Errorf("Page '%s' already exists", title)
	} else {
		return fmt.Errorf("Page name '%s' not allowed: %s", title, err.Error())
	}
}
