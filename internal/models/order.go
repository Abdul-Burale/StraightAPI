package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	Name 		string	`bson:"name"`
	Quantity	int 	`bson:"quantity"`
	Price		float64	`bson:"price"`
}

//TODO: Use Log Fatal and create Tests
//TODO: Comment with clearly

type Order struct {
	ID				primitive.ObjectID  `bson:"_id,omitempty"`
	UserID   		string			    `bson:"user_id"`
	OrderNumber		string				`bson:"order_number"`
	Items			[]Item				`bson:"items"`
	TotalAmount		float64				`bson:"total_amount"`
	Status			string				`bson:"status"` // Pending, Paid, etc.
	CreatedAt		time.Time			`bson:"created_at"`
	UpdatedAt		time.Time			`bson:"updated_at"`
}
// sumTotalPrice calculates the total price of all items in the given slice.
// Each item's price is multiplied by its quantity, and the sum of all such
// values is returned
//
// Parameters:
// 	items - A slice of Item structs, where each Item contains a Price and Quantity.
//
// Returns:
// 	A float64 representing the total price, calculated as the sum of

// 	item.Price * item.Quantity for each item in the slice
func SumTotalPrice(items []Item) float64 {
	var totalPrice float64
	for _, item := range items{
		totalPrice += item.Price * float64(item.Quantity)
	}
	return totalPrice
}
