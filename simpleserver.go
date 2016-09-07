package main

import (
    "fmt"
//    "html"
    "log"
    "net/http"
)

func main() {

//function that searches the database for malware URLs
//should search the database using something like this: GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}


// add a function to create a 2 flows (1 for malware URLs and one for ok URLs)


//this is the funtion that makes any typed URL in my local server works, and adds everything after 8081 to be pirnted at the screen
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })  
//  I can comment out his code, nothing breaks, not sure what it does
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    }) 

    log.Fatal(http.ListenAndServe(":8081", nil))

}