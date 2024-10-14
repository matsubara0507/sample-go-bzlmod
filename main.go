package main

import (
	_ "embed"
	"log"
	"net/http"
)

//go:embed index.html
var indexHtmlFile []byte

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write(indexHtmlFile)
}
