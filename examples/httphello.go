package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):] //get everything after the /hello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/shouthello/"):] //get everything after the /shouthello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/"):] //get everything after the /shouthello/ part of the URL
	fmt.Fprintf(w, "Welcome. Append /hello/{YOUR NAME} to see us say hello to you.", strings.ToUpper(remPartOfURL))
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe("localhost:9999", nil)
}
