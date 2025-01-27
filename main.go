package main

import (
	"StraightAPI/internal/api"
	"StraightAPI/internal/db"
	"net/http"
	"log"

)

	// TODO: Test Utils.LoadEnv (if Check)
	// TODO: Init Method for DB Connect

func main() {

	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	client := db.GetClient()
	defer client.Disconnect(nil)


	http.HandleFunc("/", api.HomePage) // registers the handler for this specific route
	http.HandleFunc("/Home", api.HomePage)

	// ListenAndServe automatically generates a multiplexer
	// mux handles the logic of separating routes with a function called ServeHTTP
	// Consider a route as a key in a dict
	// and the handler as its value
	// Mux[Route-> '/home'] = homepage()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
