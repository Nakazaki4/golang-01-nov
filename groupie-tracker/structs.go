package main

// Band represents an artist
type Band struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// Locations represents the locations JSON
type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// ConcertDates represents the concert dates JSON
type ConcertDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// DatesLocations represents the dates-locations JSON
type DatesLocations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}