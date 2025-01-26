package models

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Item struct {
	Name 		string	`json:"name"`
	Quantity	int 	`json:"quantity"`
	Price		float64	`json:"price"`
}

type Order struct {
	ID				primitive.ObjectID  `json:"id"`
	UserID   		string			    `json:"user_id"`
	OrderNumber		string				`json:"order_number"`
	Items			[]Item				`json:"items"`
	TotalAmount		float64				`json:"total_amount"`
	Status			string				`json:"status"` // Pending, Paid, etc.
	CreatedAt		time.Time			`json:"created_at"`
	UpdatedAt		time.Time			`json:"updated_at"`
}

func CreateOrder(client *mongo.Client, order *Order) (*Order, error) {
	collection := client.Database("StraightApiDB").Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	// Insert order into the collection
	res, err := collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}

	order.ID =res.InsertedID.(primitive.ObjectID)
	return order, nil
}
