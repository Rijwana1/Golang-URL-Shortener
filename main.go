package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"poc/url-shortener/model"
	"poc/url-shortener/store"

	"net/http"

	"github.com/gorilla/mux"
)

func homePage() http.Handler {
	return http.FileServer(http.Dir("./public/"))
}

func generateShortURL() string {
	b := make([]byte, 2)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	requestedURL := r.URL.Query().Get("url")
	shortURL := generateShortURL()

	URL := *&model.URLTable{ShortURL: shortURL, LongURL: requestedURL}

	err := store.Create(URL)
	if err != nil {
		log.Printf("Internal server error %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Unable to create,something went to worng"))
	} else {
		fmt.Print("its work correctly")
		w.WriteHeader(200)
		w.Write([]byte("Created New Record successfully"))
	}
	fmt.Fprintf(w, "http://localhost:10000/s/"+URL.ShortURL)

}

func redirectToActualURL(w http.ResponseWriter, r *http.Request) {

	shortURL := mux.Vars(r)["shorturl"]

	out, err := store.Find(shortURL)
	if err != nil {
		log.Printf("Internal server error %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Unable to get data,something went to worng"))
	} else {
		fmt.Print("it works correctly")
		w.WriteHeader(200)
	}
	fmt.Print(out.ShortURL, shortURL)

	if out.ShortURL == shortURL {
		http.Redirect(w, r, out.LongURL, 301)
		return
	}

	http.Error(w, "Not found", http.StatusNotFound)
}

func handleServer() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/s/{shorturl}", redirectToActualURL).Methods("GET")
	router.HandleFunc("/shorten-url", shortenURL).Methods("GET")
	router.PathPrefix("/").Handler(homePage())
	fmt.Println("localhost")
	log.Fatal(http.ListenAndServe(":10000", router))

}

func main() {
	handleServer()
}
