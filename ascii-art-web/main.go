package main

import (
	asciiart "HttpTraining/ascii-art"
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

type Data struct {
	Error       string
	AsciiResult string
}

func main() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.ListenAndServe(":8080", nil)
}

// Method to show the home page to which the user can interact
func rootHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if request.URL.Path != "/"{
		
		http.Error(writer, "", http.StatusNotFound)
		return
	}

	err := tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
	}
}

func asciiArtHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not alloWed,", http.StatusMethodNotAllowed)
		return
	}

	text := request.FormValue("text")
	banner := request.FormValue("banner")

	error := asciiart.TextToDrawValidation(text)
	if error != nil {
		unsucsseful := Data{
			Error:       error.Error(),
			AsciiResult: "",
		}
		tmpl.Execute(writer, unsucsseful)
		return
	}

	asciiArtResult := asciiart.GenerateAscii(text, banner)
	succesful := Data{
		Error:       "",
		AsciiResult: asciiArtResult,
	}

	tmpl.Execute(writer, succesful)

	fmt.Printf("text is:%s banner is:%s\n", text, banner)
}
