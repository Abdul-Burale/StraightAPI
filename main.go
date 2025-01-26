package main

import (
	"StraightAPI/internal/models"
	"StraightAPI/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	"net/http"
	"log"
	"time"

	// TODO: Not in go ROUTE need to ADD
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

	utils.LoadEnv()

	mongoURI := utils.GetMongoURI()

	client, err := mongo.Connect(nil, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	defer client.Disconnect(nil)

	order := models.Order{
		UserID:				"user1",
		OrderNumber:		"0",
		Items:	[]models.Item{
			{Name: "Burger", Quantity: 2, 	Price: 5.99},
			{Name: "Fries",  Quantity: 1,	Price: 2.99},
		},
		TotalAmount: 		14.97,
		Status:				"pending",
		CreatedAt:			time.Now(),
		UpdatedAt:			time.Now(),
	}

	createdOrder, err := models.CreateOrder(client, &order)
	if err != nil {
	log.Fatalf("Error creating order: %v", err)
	}

	log.Printf("Created order: %+v,", createdOrder)

	http.HandleFunc("/", homePage) // registers the handler for this specific route


	// ListenAndServe automatically generates a multiplexer
	// mux handles the logic of separating routes with a function called ServeHTTP
	// Consider a route as a key in a dict
	// and the handler as its value
	// Mux[Route-> '/home'] = homepage()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
