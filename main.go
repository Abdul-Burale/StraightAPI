package main

import (
	"fmt"
	"net/http"
	"log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the landing page")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
	//	handleRequests()
func main() {
    http.HandleFunc("/", homePage) // registers the handler for this specific route

	// ListenAndServe automatically generates a multiplexer
	// mux handles the logic of separating routes with a function called ServeHTTP
	// Consider a route as a key in a dict
	// and the handler as its value
	// Mux[Route-> '/home'] = homepage()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
