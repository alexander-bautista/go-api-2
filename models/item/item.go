package item

import (
	"context"

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
		cursor.Decode(&oneItem)
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
