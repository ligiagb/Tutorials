package main

import (
    "fmt"
    "html"
    "io"
    "log"
    "net/http"
    "unicode/utf8"
)
//json to tell the client website is safe and return the same URL - I probably have to save the URL requested somewhere
//took from https://tutorialedge.net/creating-simple-restful-json-api-with-go and just adding jsons doesn't break it
//TO DO:figure out how routers work and why I can't install it
type ValidationResponce struct {
	OriginalURL string `json:"url"`    
	SafeURL      bool   `json:"safety"` //not sure if should go for the safe or unsafe one here
	
}
//function that searches the database for malware URLs
//should search the database using something like this: GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
func DataBaseValidation(url string) bool {
    // if function to consider all websites with an odd number of character as malware
   urlLength := utf8.RuneCountInString(html.EscapeString(url))
    if urlLength%2==0 {
        return true
    } else {
        return false
    }

}
//  function with 2 flows depending on the URL requested (one for malware URLs and one for ok URLs)
func urlinfo(w io.Writer,url string) {
    validUrl := DataBaseValidation(url)
    if(validUrl == true) {
         fmt.Fprintf(w,"URL IS GOOD")
    } else {
        fmt.Fprintf(w,"URL IS BAD")
    }
}
func main() {

//this is handling the urlinfo function to handle url requests
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        urlinfo(w, html.EscapeString(r.URL.Path))
    })         

    log.Fatal(http.ListenAndServe(":8081", nil))

}