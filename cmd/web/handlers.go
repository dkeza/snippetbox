package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		rootDir + "ui/html/home.page.tmpl",
		rootDir + "ui/html/base.layout.tmpl",
		rootDir + "ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Interval Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Interval Server Error", 500)
	}

}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id")
	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Show snippet with id %v", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create snippet"))
}
