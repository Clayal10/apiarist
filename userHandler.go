package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Clayal10/mathGen/gen"
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
		var err error

		// Parse the submission form if it is a Post
		if err := request.ParseForm(); err != nil {
			fmt.Println("Could not parse form submission")
			return
		}

		var iterations int64
		var inertia float64
		var cogCoef float64
		var socCoef float64
		// Grab the input text box ID
		iterationsString := request.FormValue("iterations")
		inertiaString := request.FormValue("inertia")
		cogCoefString := request.FormValue("cog-coef")
		socCoefString := request.FormValue("soc-coef")

		if iterations, err = strconv.ParseInt(iterationsString, 10, 64); err != nil {
			fmt.Println(err)
			return
		}

		if inertia, err = strconv.ParseFloat(inertiaString, 64); err != nil {
			fmt.Println(err)
			return
		}

		if cogCoef, err = strconv.ParseFloat(cogCoefString, 64); err != nil {
			fmt.Println(err)
			return
		}

		if socCoef, err = strconv.ParseFloat(socCoefString, 64); err != nil {
			fmt.Println(err)
			return
		}

		dataInput := gen.UserInput{
			Iterations: int(iterations),
			Inertia:    inertia,
			CogCoef:    cogCoef,
			SocCoef:    socCoef,
		}

		go gen.PSOSineGen(dataInput)

		/*
			template, err := template.ParseFiles("./template/output.html")
			if err != nil {
				fmt.Println("Could not parse template")
				fmt.Println(err)
				return
			}

			// Execute the template
			if err = template.Execute(write, nil); err != nil {
				fmt.Println("Could not execute template")
				fmt.Println(err)
				return
			}
		*/
	}
}

func graphHandler(write http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		fmt.Println("Graphing!!")
	}
}
