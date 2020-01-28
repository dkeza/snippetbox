package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	rootDir = "../../"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir(rootDir + "ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %v \n", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
