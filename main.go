package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type URLTable struct {
	ShortURL string
	LongURL  string
}

var urls []URLTable

func homePage() http.Handler {
	return http.FileServer(http.Dir("./public/"))
}

func generateShortURL() string {
	b := make([]byte, 3)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	requestedURL := r.URL.Query().Get("url")
	shortURL := generateShortURL()

	urls = append(urls, URLTable{ShortURL: shortURL, LongURL: requestedURL})
	fmt.Fprintf(w, "http://localhost:10000/s/"+shortURL)
}

func redirectToActualURL(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shorturl"]

	for _, url := range urls {
		if url.ShortURL == shortURL {
			http.Redirect(w, r, url.LongURL, http.StatusMovedPermanently)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

func handleServer() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/s/{shorturl}", redirectToActualURL).Methods("GET")
	router.HandleFunc("/shorten-url", shortenURL).Methods("GET")
	router.PathPrefix("/").Handler(homePage())

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleServer()
}
