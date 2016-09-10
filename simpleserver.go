package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/gorilla/mux"
)

//json to tell the client website is safe and return the same URL - I probably have to save the URL requested somewhere
//took from https://tutorialedge.net/creating-simple-restful-json-api-with-go and just added jsons doesn't break it
type ValidationResponce struct {
	OriginalURL string `json:"url"`
	FlaggedURL  bool   `json:"flagged"` //not sure if should go for the safe or unsafe one here

}

//function that searches the database for malware URLs
//(from Skye) should search the database using something like this: GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
func DatabaseValidation(url string) bool {
	// if function to consider all websites with an odd number of character as malware for now
	urlLength := utf8.RuneCountInString(html.EscapeString(url))
	if urlLength%2 == 0 {
		return true
	} else {
		return false
	}
}

//function to build the json response
func JsonResponse(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	UnsafeUrl := DatabaseValidation(url)
	Response := ValidationResponce{OriginalURL: url, FlaggedURL: UnsafeUrl}
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		panic(err)
	}
}

//  function with 2 flows depending on the URL requested (one for malware URLs and one for ok URLs)
func urlinfo(w io.Writer, url string) {
	UnsafeUrl := DatabaseValidation(url)
	if UnsafeUrl == true {
		fmt.Fprintf(w, "URL IS BAD, please try another URL")
	} else {
		fmt.Fprintf(w, "URL IS GOOD")
	}

}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	//this is handling the urlinfo function to handle url requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlinfo(w, html.EscapeString(r.URL.Path))
	})
	// json response is here, not really sure hot to make it print on borwnser -  have to figure out how to cache it
	router.HandleFunc("/urlinfo/1/{urlinfo}", JsonResponse).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
