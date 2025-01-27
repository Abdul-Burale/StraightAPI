package db

import (
	"StraightAPI/internal/models"
	"StraightAPI/internal/utils"
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DB struct {
	Client  			*mongo.Client
	Database			*mongo.Database
	OrdersCollection	*mongo.Collection
}

var globalDB *DB

func InitDB() error {

	utils.LoadEnv()

	mongoURI	 	:= utils.GetMongoURI()
	dbName      	:= utils.GetMongoDatabaseName()
	collectionName  := utils.GetMongoDatabaseCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))


	if err != nil {
		return err
	}


	database := client.Database(dbName)
	ordersCollection := database.Collection(collectionName)

	globalDB = &DB{
		Client: 			 client,
		Database:			 database,
		OrdersCollection:	 ordersCollection,
	}
	return nil
}

func initDBContext() ( context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return ctx, cancel
}

func GetClient() *mongo.Client {
	return globalDB.Client
}

func CreateOrder(order *models.Order) (*models.Order, error) {
	ctx, cancel := initDBContext()
	defer cancel()

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.TotalAmount = models.SumTotalPrice(order.Items)

	collection := globalDB.OrdersCollection

	// Insert order into the collection
	res, err := collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}

	order.ID = res.InsertedID.(primitive.ObjectID)
	return order, nil
}


func UpdateOrder(order *models.Order) (*models.Order, error) {
	ctx, cancel := initDBContext()
	defer cancel()

	// The filter to find the document we want to update -> using the order ID
	filter := bson.M{
		"_id": order.ID,
	}

	order.TotalAmount = models.SumTotalPrice(order.Items)
	order.UpdatedAt = time.Now()
	// the update operation
	update := bson.M{
		"$set": bson.M{
			"items": order.Items,
			"total_amount": order.TotalAmount,
			"updated_at": order.UpdatedAt,
		},
	}


	collection := globalDB.OrdersCollection
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

func DeletedOrder(order *models.Order) (*models.Order, error) {
	ctx, cancel := initDBContext()
	defer cancel()

	filter := bson.M{"_id": order.ID,}

	collection := globalDB.OrdersCollection
	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("no document found with the given_id")

	}
	return order, nil
}
