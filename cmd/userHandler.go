package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Input will be which math function they want to use and the value to input
type UserInput struct {
	Function string
	InputVal float64
}

func mainPageHandler(write http.ResponseWriter, request *http.Request) {
	//Reads content of html file and returns a template
	template, err := template.ParseFiles("./web/home.html")
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

func submitHandler(write http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		// Parse the submission form if it is a Post
		err := request.ParseForm()
		if err != nil {
			fmt.Println("Could not parse form submission")
			return
		}

		// Grab the input text box ID
		inputValBuffer := request.FormValue("inputVal")
		inputVal, err := strconv.ParseFloat(inputValBuffer, 64)
		if err != nil {
			fmt.Println("Could not convert to float")
			return
		}

		data := UserInput{
			Function: "Default",
			InputVal: inputVal,
		}

		// The template for /submit is also the home template for now
		template, err := template.ParseFiles("./web/output.html")
		if err != nil {
			fmt.Println("Could not parse template")
			return
		}

		// Execute the template
		err = template.Execute(write, data)
		if err != nil {
			fmt.Println("Could not execute template")
			return
		}
	}
}
