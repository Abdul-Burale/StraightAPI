package main

import (
	"StraightAPI/internal/api"
	"StraightAPI/internal/db"
	"StraightAPI/internal/middleware"
	"StraightAPI/internal/auth"
	"net/http"
	"log"
	//	"context"
)

	// TODO: Test Utils.LoadEnv (if Check)
	// TODO: Init Method for DB Connect

func main() {

	// Init Database
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	client := db.GetClient()
	defer client.Disconnect(nil)


	// Init Firebase
	if err := auth.InitFireBase(); err != nil {
		log.Fatalf("Error initalizing Firebase: %v", err)
	}


	// End Points
	http.HandleFunc("/", middleware.CorsMiddleware(api.HomePage)) // registers the handler for this specific route
	http.HandleFunc("OPTIONS /", middleware.CorsMiddleware(func(w http.ResponseWriter, r *http.Request) {}))

	http.HandleFunc("GET /Home", middleware.CorsMiddleware(middleware.AuthMiddleware(api.HomePage)))
	http.HandleFunc("OPTIONS /Home", middleware.CorsMiddleware(func(w http.ResponseWriter, r *http.Request) {}))

	http.HandleFunc("GET /check", middleware.CorsMiddleware(middleware.AuthMiddleware(api.SuccessAuth)))
	http.HandleFunc("OPTIONS /check", middleware.CorsMiddleware(func (w http.ResponseWriter, r *http.Request) {}))


	// ListenAndServe automatically generates a multiplexer
	// mux handles the logic of separating routes with a function called ServeHTTP
	// Consider a route as a key in a dict
	// and the handler as its value
	// Mux[Route-> '/home'] = homepage()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
