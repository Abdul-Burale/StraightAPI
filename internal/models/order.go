package models

import (
	"fmt"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
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

func initDBContext(client *mongo.Client) (*mongo.Collection, context.Context, context.CancelFunc) {
	collection := client.Database("StraightApiDB").Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return collection, ctx, cancel
}


func CreateOrder(client *mongo.Client, order *Order) (*Order, error) {
	collection, ctx, cancel := initDBContext(client)
	defer cancel()

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.TotalAmount = sumTotalPrice(order.Items)


	// Insert order into the collection
	res, err := collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}

	order.ID = res.InsertedID.(primitive.ObjectID)
	return order, nil
}


func UpdateOrder(client *mongo.Client, order *Order) (*Order, error) {
	collection, ctx, cancel := initDBContext(client)
	defer cancel()

	// The filter to find the document we want to update -> using the order ID
	filter := bson.M{
		"_id": order.ID,
	}

	order.TotalAmount = sumTotalPrice(order.Items)
	order.UpdatedAt = time.Now()
	// the update operation
	update := bson.M{
		"$set": bson.M{
			"items": order.Items,
			"total_amount": order.TotalAmount,
			"updated_at": order.UpdatedAt,
		},
	}


	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// TODO:  Use Log FatalF
	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("no document found to update")
	}

	return order, nil
}

func DeletedOrder(client *mongo.Client, order *Order) (*Order, error) {
	collection, ctx, cancel := initDBContext(client)
	defer cancel()

	filter := bson.M{"_id": order.ID,}

	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("no document found with the given_id")
	}

	return order, nil
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

func sumTotalPrice(items []Item) float64 {
	var totalPrice float64
	for _, item := range items{
		totalPrice += item.Price * float64(item.Quantity)
	}
	return totalPrice
}
