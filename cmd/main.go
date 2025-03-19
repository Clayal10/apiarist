package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home/", mainPageHandler)
	http.HandleFunc("/submit/", submitHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
