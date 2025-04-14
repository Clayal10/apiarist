package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/submit/", submitHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":8200", nil)
	if err != nil {
		log.Fatal(err)
	}
}
