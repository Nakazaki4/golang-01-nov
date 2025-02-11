package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var tmpl *template.Template

type Err struct {
	Message string
	Value   int
}

// Handler function to fetch and display artists
func artistHandler(w http.ResponseWriter, r *http.Request) {
	if !ValidatePath(w, r) {
		fmt.Println("Invalid path")
		return
	}
	if !ValidateHttpMethod(w, http.MethodGet, r.Method) {
		return
	}
	apiURL := "https://groupietrackers.herokuapp.com/api/artists" // API URL
	var bands []Band

	// Fetch all artists
	err := FetchJSON(apiURL, &bands)
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	renderTemplate(w, bands, http.StatusOK, "homePage.html")
}

// Handler to display a single artist's details
func artistDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ArtistDetailHAndler")
	if !ValidatePath(w, r) {
		fmt.Println("Invalid path")
		return
	}
	if !ValidateHttpMethod(w, http.MethodGet, r.Method) {
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing artist ID", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id)
	locationsApiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%s", id) // API URL
	datesApiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%s", id)
	relationsApiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)

	var artist Band
	err := FetchJSON(apiURL, &artist)
	if err != nil {
		http.Error(w, "Failed to fetch artist details", http.StatusInternalServerError)
		return
	}

	var locations Locations
	err = FetchJSON(locationsApiURL, &locations)
	if err != nil {
		http.Error(w, "Failed to fetch Location", http.StatusInternalServerError)
		return
	}

	var concertDates ConcertDates
	err = FetchJSON(datesApiURL, &concertDates)
	if err != nil {
		http.Error(w, "Failed to fetch ConcertDates", http.StatusInternalServerError)
		return
	}

	var datesLocations DatesLocations
	err = FetchJSON(relationsApiURL, &datesLocations)
	if err != nil {
		http.Error(w, "Failed to fetch DatesLocations", http.StatusInternalServerError)
		return
	}

	data := struct {
		Artist         Band
		Location       Locations
		ConcertDate    ConcertDates
		DatesLocations DatesLocations
	}{
		Artist:         artist,
		Location:       locations,
		ConcertDate:    concertDates,
		DatesLocations: datesLocations,
	}

	renderTemplate(w, data, http.StatusOK, "artistDetails.html")
}

func ValidateHttpMethod(w http.ResponseWriter, expectedMethod, actualMethod string) bool {
	if expectedMethod != actualMethod {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func ValidatePath(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != "/" && r.URL.Path != "/artist/" {
		return false
	}
	return true
}

func notFoundHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	Err := &Err{
		Message: http.StatusText(http.StatusNotFound),
		Value:   http.StatusNotFound,
	}
	tmpl := template.Must(template.ParseFiles("./templates/error.html"))
	tmpl.Execute(w, Err)
}

func renderTemplate(w http.ResponseWriter, data interface{}, statusCode int, template string) {
	if isFileAvailable(template) {
		w.WriteHeader(statusCode)
		tmpl.ParseFiles("templates/" + template)
		tmpl.Execute(w, data)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "500 - Internal server error")
	}
}

func isFileAvailable(file string) bool {
	_, err := os.Stat("./templates/" + file)
	return err == nil
}
