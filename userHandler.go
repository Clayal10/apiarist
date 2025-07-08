package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/Clayal10/mathGen/gen"
)

var swarm *gen.Swarm

// Creates home page template.
func mainPageHandler(write http.ResponseWriter, request *http.Request) {
	//Reads content of html file and returns a template
	template, err := template.ParseFiles("./template/home.html")
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

		dataInput := &gen.UserInput{}
		if err = json.Unmarshal(body, dataInput); err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			return
		}

		write.WriteHeader(http.StatusOK)

		swarm = &gen.Swarm{}
		swarm.InitSwarm(dataInput)
		go gen.PSOSineGen(swarm, dataInput)
	}
}

type GraphData struct {
	Data []float64
}

func graphHandler(write http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		if swarm == nil {
			write.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Swarm not initialized.")
			return
		}
		fmt.Println("Starting To Generate Network")

		data := GraphData{
			Data: swarm.GetValues(),
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
