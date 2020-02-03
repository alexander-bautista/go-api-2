package comics

import (
	"context"
	"log"
	"math/rand"

	"github.com/alexander-bautista/go-api-2/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Collection = "comics"
)

// Comic :  comic model
type Comic struct {
	Id       int     `json:"id" bson:"id,omitempty"`
	Title    string  `json:"title,omitempty"`
	Isbn     string  `json:"isbn,omitempty"`
	Format   string  `json:"format,omitempty"`
	Dates    []date  `json:"dates"`
	Prices   []price `json:"prices"`
	Quantity int     `json:"quantity"`
}

type date struct {
	Type string `json:"type"`
	Date string `json:"date"`
}

type price struct {
	Type  string  `json:"type"`
	Price float32 `json:"price"`
}

// GetAll : get all items
func GetAll() (items []Comic) {
	collection := db.Client.Database("todo").Collection(Collection)
	cursor, _ := collection.Find(context.TODO(), bson.M{})

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var oneItem Comic
		err := cursor.Decode(&oneItem)

		if err != nil {
			log.Fatal(err)
		}

		items = append(items, oneItem)
	}
	return
}

// GetOne : Get one item
func GetOne(id int) (item Comic) {
	collection := db.GetClient().Database("todo").Collection(Collection)
	collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&item)
	return
}

// Add : Add one item to collection
func Add(comic Comic) interface{} {
	collection := db.Client.Database("todo").Collection(Collection)

	result, err := collection.InsertOne(context.TODO(), comic)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

// Add : Add one item to collection
func AddMany(items []interface{}) interface{} {
	collection := db.Client.Database("todo").Collection(Collection)

	result, err := collection.InsertMany(context.TODO(), items)

	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedIDs
}

// FindOneAndUpdate : finds one comic and update if exist or create it otherwise
func FindOneAndUpdate(comic Comic) (bson.M, error) {
	collection := db.Client.Database("todo").Collection(Collection)
	filter := bson.M{"id": comic.Id}

	// Create an instance of an options and set the desired options
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	// Set quantity random
	comic.Quantity = rand.Intn(1000)

	result := collection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": comic}, &opt)

	if result.Err() != nil {
		return nil, result.Err()
	}
	item := bson.M{}

	decodeErr := result.Decode(&item)

	return item, decodeErr
}
