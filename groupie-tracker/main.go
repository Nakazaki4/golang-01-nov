package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", artistHandler)
	mux.HandleFunc("/artist/", artistDetailHandler)

	serve := &http.Server{
		Addr:    ":8080",
		Handler: logMiddleware(notFoundMiddleware(mux)),
	}

	fmt.Println("Server started on http://localhost:8080")
	if err := serve.ListenAndServe(); err != nil {
		log.Fatalf("500 - Internal Server Error: %v", err)
	}
}

func notFoundMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedPaths := map[string]bool{
			"/":                             true,
			"/artist/":                      true,
			"/static/css/homePage.css":      true,
			"/static/css/errorPage.css":     true,
			"/static/css/artistDetails.css": true,
		}
		if !allowedPaths[r.URL.Path] {
			notFoundHandler(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// FetchJSON makes an HTTP GET request and unmarshals the response
func FetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// If Unmarshalling fails an error will be returned
	return json.Unmarshal(body, target)
}
