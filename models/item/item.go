package item

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Item : represents a single database item
type Item struct {
	ID     int    `json:"id,omitempty bson:"id,omitempty"`
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty"`
}

// GetAll : get all items
func GetAll(client *mongo.Client) (items []Item) {
	collection := client.Database("todo").Collection("items")
	cursor, _ := collection.Find(context.TODO(), bson.M{})

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var oneItem Item
		err := cursor.Decode(&oneItem)

		if err != nil {
			log.Fatal(err) // ????
		}

		items = append(items, oneItem)
	}
	return
}

// GetOne : Ge one item
func GetOne(client *mongo.Client, id int) (item Item) {
	collection := client.Database("todo").Collection("items")
	collection.FindOne(context.TODO(), Item{ID: id}).Decode(&item)
	return
}

// Add : Add one item to collection
func Add(client *mongo.Client, item Item) interface{} {
	collection := client.Database("todo").Collection("items")

	result, err := collection.InsertOne(context.TODO(), item)

	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

// RemoveOne : Remove one item by id
func RemoveOne(client *mongo.Client, id int) int64 {
	collection := client.Database("todo").Collection("items")

	result, err := collection.DeleteOne(context.TODO(), Item{ID: id})

	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}
