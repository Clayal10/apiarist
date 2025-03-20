package main

import (
	"fmt"
	network "github.com/Clayal10/mathGen/pkg"
	"html/template"
	"net/http"
	"strconv"
)

// Input will be which math function they want to use and the value to input
type UserInput struct {
	Function string
	InputVal float64
}

type UserOutput struct {
	Function  string
	OutputVal float64
}

// Creates home page template.
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
		newVal, err := strconv.ParseFloat(inputValBuffer, 64)
		if err != nil {
			fmt.Println("Could not convert to float")
			return
		}

		// Create an output struct after parsing the user input
		data := UserOutput{
			Function:  inputFuncBuffer,
			OutputVal: network.SineGen(newVal),
		}

		// The template for /submit is also the home template for now
		template, err := template.ParseFiles("./web/output.html")
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
