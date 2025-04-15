package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Clayal10/mathGen/lib/parser"
)

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
		// Parse the submission form if it is a Post
		if err := request.ParseForm(); err != nil {
			fmt.Println("Could not parse form submission")
			return
		}

		// Grab the input text box ID
		inputFuncBuffer := request.FormValue("functionVal")
		inputValBuffer := request.FormValue("inputVal")
		inputLearningBuffer := request.FormValue("learning")

		newVal, err := strconv.ParseFloat(inputValBuffer, 64)
		if err != nil {
			fmt.Println("Could not convert to float")
			return
		}

		dataInput := parser.UserInput{
			Function: inputFuncBuffer,
			InputVal: newVal,
			Learning: inputLearningBuffer,
		}

		// Create an output struct after parsing the user input
		data := parser.TakeUserInput(dataInput)

		// The template for /submit is also the home template for now
		template, err := template.ParseFiles("./template/output.html")
		if err != nil {
			fmt.Println("Could not parse template")
			return
		}

		// Execute the template
		if err = template.Execute(write, data); err != nil {
			fmt.Println("Could not execute template")
			return
		}
	}
}
