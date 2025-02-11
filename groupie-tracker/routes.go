package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl *template.Template

type Err struct {
	Message string
	Value   int
}

type APIEndPoint struct {
	URL    string
	Target interface{}
}

func init() {
	var err error
	tmpl, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatalf("Failed to initialize templates: %v", err)
	}
}

// Handler function to fetch and display artists in home page
func artistHandler(wr http.ResponseWriter, r *http.Request) {
	if !ValidatePath(r) {
		fmt.Println("Invalid path")
		return
	}
	if !ValidateHttpMethod(wr, http.MethodGet, r.Method) {
		return
	}
	apiURL := "https://groupietrackers.herokuapp.com/api/artists" // API URL
	var bands []Band

	// Fetch all artists
	err := FetchJSON(apiURL, &bands)
	if err != nil {
		renderError(wr, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	renderTemplate(wr, bands, "homePage.html")
}

// Handler to display a single artist's details
func artistDetailHandler(wr http.ResponseWriter, r *http.Request) {
	if !ValidatePath(r) || !ValidateHttpMethod(wr, http.MethodGet, r.Method) {
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		renderError(wr, http.StatusBadRequest, "Missing artist ID")
		return
	}

	baseUrl := "https://groupietrackers.herokuapp.com/api"

	data := struct {
		Artist         Band
		Location       Locations
		ConcertDate    ConcertDates
		DatesLocations DatesLocations
	}{}

	endPoints := []APIEndPoint{
		{fmt.Sprintf("%s/artists/%s", baseUrl, id), &data.Artist},
		{fmt.Sprintf("%s/locations/%s", baseUrl, id), &data.Location},
		{fmt.Sprintf("%s/dates/%s", baseUrl, id), &data.ConcertDate},
		{fmt.Sprintf("%s/relation/%s", baseUrl, id), &data.DatesLocations},
	}

	for _, endpoint := range endPoints {
		if err := FetchJSON(endpoint.URL, endpoint.Target); err != nil {
			renderError(wr, http.StatusInternalServerError, "Failed to fetch data")
			return
		}
	}

	renderTemplate(wr, data, "artistDetails.html")
}

func ValidateHttpMethod(wr http.ResponseWriter, expectedMethod, actualMethod string) bool {
	if expectedMethod != actualMethod {
		renderError(wr, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return false
	}
	return true
}

func ValidatePath(r *http.Request) bool {
	validPaths := map[string]bool{
		"/":        true,
		"/artist/": true,
	}
	return validPaths[r.URL.Path]
}

func notFoundHandler(wr http.ResponseWriter) {
	renderError(wr, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func renderTemplate(wr http.ResponseWriter, data interface{}, template string) {
	if isFileAvailable(template) {
		err := tmpl.ExecuteTemplate(wr, template, data)
		if err != nil {
			renderError(wr, http.StatusInternalServerError, "Template Rendering Error")
		}
	} else {
		renderError(wr, http.StatusInternalServerError, template+" not available")
	}
}

func renderError(wr http.ResponseWriter, statusCode int, msg string) {
	if isFileAvailable("errorPage.html") {
		wr.WriteHeader(statusCode)
		renderTemplate(wr, &Err{Message: msg, Value: statusCode}, "errorPage.html")
	} else {
		fallbackErrorMessage(wr)
	}
}

func fallbackErrorMessage(wr http.ResponseWriter) {
	wr.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(wr, "Error 500: Website is under maintenance for security issues.")
}

func isFileAvailable(file string) bool {
	_, err := os.Stat("./templates/" + file)
	return err == nil
}
