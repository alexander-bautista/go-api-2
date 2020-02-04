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
	col = "comics"
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

// getAll : get all items
func getAll() (items []Comic) {

	ctx, client := db.Connect()

	collection := client.Database("todo").Collection(col)
	cursor, _ := collection.Find(ctx, bson.M{})

	defer func() {
		cursor.Close(ctx)
		client.Disconnect(ctx)
	}()

	for cursor.Next(ctx) {
		var oneItem Comic
		err := cursor.Decode(&oneItem)

		if err != nil {
			log.Fatal(err)
		}

		items = append(items, oneItem)
	}
	return
}

// getOne : Get one item
func getOne(id int) (item Comic) {

	ctx, client := db.Connect()

	defer client.Disconnect(ctx)

	collection := client.Database("todo").Collection(col)
	collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&item)
	return
}

// Add : Add one item to collection
/*func Add(comic Comic) interface{} {
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
*/
// findOneAndUpdate : finds one comic and update if exist or create it otherwise
func findOneAndUpdate(comic Comic) (bson.M, error) {

	_, client := db.Connect()

	defer client.Disconnect(context.Background())

	collection := client.Database(db.DbName).Collection(col)
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

	result := collection.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": comic}, &opt)

	if result.Err() != nil {
		return nil, result.Err()
	}
	item := bson.M{}

	decodeErr := result.Decode(&item)

	return item, decodeErr
}

// FindOneAndUpdate : finds one comic and update if exist or create it otherwise
func findAndUpdateMany(comics []Comic) ([]interface{}, error) {

	ctx, client := db.Connect()

	defer client.Disconnect(ctx)

	collection := client.Database(db.DbName).Collection(col)

	// Create an instance of an options and set the desired options
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	var items = make([]interface{}, len(comics))

	for i, comic := range comics {
		// Set quantity random
		comic.Quantity = rand.Intn(1000)

		filter := bson.M{"id": comic.Id}

		result := collection.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": comic}, &opt)

		if result.Err() != nil {
			return nil, result.Err()
		}

		item := bson.M{}
		decodeErr := result.Decode(&item)

		if decodeErr != nil {
			return nil, decodeErr
		}

		items[i] = item
	}

	return items, nil
}
