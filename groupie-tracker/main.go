package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Band struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"` // link
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsURL string   `json:"locations"`    // link
	ConcertDates string   `json:"concertDates"` // link
	Relations    string   `json:"relations"`    // link
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     []string `json:"dates"`
}

var dates []string

func main() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	bands := []Band{}
	json.Unmarshal(content, &bands)

	for i := 0; i < len(bands); i++ {
		fetchLocations(bands[i].LocationsURL)
		fmt.Println(bands[i].Name)
	}
}

func fetchLocations(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	var location Locations
	json.Unmarshal(content, &location)

	for i := 0; i < len(location.Dates); i++ {
		fetchDates(location.Dates[i])
		fmt.Print(location.Locations)
	}
}

func fetchDates(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(content)
	json.Unmarshal(content, &dates)
	fmt.Println(dates)
}
