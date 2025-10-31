package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/Clayal10/mathGen/swarm"
)

func main() {
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/submit/", submitHandler)
	http.HandleFunc("/graph/", graphHandler)
	http.HandleFunc("/stop/", stopHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":8200", nil)
	if err != nil {
		log.Fatal(err)
	}
}

var s *swarm.Swarm

// Creates home page template.
func mainPageHandler(write http.ResponseWriter, request *http.Request) {
	//Reads content of html file and returns a template
	template, err := template.ParseFiles("./static/html/home.html")
	if err != nil {
		fmt.Println("Could not parse template")
		return
	}

	err = template.Execute(write, nil)
	if err != nil {
		fmt.Println("Could not execute template")
		return
	}
}

// Takes user submission and prepares it for output on a new page
func submitHandler(write http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var body []byte
		var err error

		if body, err = io.ReadAll(request.Body); err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			return
		}

		dataInput := &swarm.UserInput{}
		if err = json.Unmarshal(body, dataInput); err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			return
		}

		write.WriteHeader(http.StatusOK)

		s = &swarm.Swarm{}
		s.InitSwarm(dataInput)
		go s.PSOSineGen()
	}
}

type GraphData struct {
	Data []float64
}

func graphHandler(write http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		if s == nil {
			write.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Swarm not initialized.")
			return
		}

		data := GraphData{
			Data: s.GetValues(),
		}
		dataJSON, err := json.Marshal(data)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Could not marshal data for JSON")
			return
		}
		write.Write(dataJSON)
	}
}

func stopHandler(write http.ResponseWriter, request *http.Request) {
	if s == nil {
		write.WriteHeader(http.StatusInternalServerError)
	}
	s.Stop()
	write.WriteHeader(http.StatusOK)
}
