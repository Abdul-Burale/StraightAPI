package api

import (
	//"StraightAPI/internal/db"
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to landing page")
	fmt.Println("Endpoint Hit: homePage")
}

func SuccessAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Successful Auth")
	fmt.Println("Endpoint hit: SucessAuth")
}

/*
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieved request to create order")

	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Printf("Error decoding request body %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	//	createdOrder, err = models.CreateOrder()

}
*/
