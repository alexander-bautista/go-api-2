package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	data "github.com/alexander-bautista/go-api-2/data"
	"github.com/alexander-bautista/go-api-2/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	col    = "comics"
	dbName = "todo"
)

type ComicRepositoryImpl struct {
	collection *mongo.Collection
}

// NewComicRepositoryImpl will create an object that represent the article.Repository interface
func NewComicRepositoryImpl(col *mongo.Collection) data.ComicRepository {
	return &ComicRepositoryImpl{col}
}

// getOne : Get one item
func (r *ComicRepositoryImpl) GetOne(id int) (comic *model.Comic, err error) {

	ctx, client := connect()

	defer client.Disconnect(ctx)

	collection := client.Database(dbName).Collection(col)
	err = collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&comic)

	if err != nil {
		return nil, err
	}

	return comic, err
}

// getAll : get all items
func (r *ComicRepositoryImpl) GetAll() ([]*model.Comic, error) {

	ctx, client := connect()

	collection := client.Database(dbName).Collection(col)
	cursor, _ := collection.Find(ctx, bson.M{})

	defer func() {
		cursor.Close(ctx)
		client.Disconnect(ctx)
	}()

	items := make([]*model.Comic, 0)

	for cursor.Next(ctx) {
		oneItem := &model.Comic{}
		err := cursor.Decode(&oneItem)

		if err != nil {
			return nil, err
		}

		items = append(items, oneItem)
	}

	return items, nil
}

func connect() (context.Context, *mongo.Client) {
	connectionString := "mongodb+srv://todo_user:todo2020@traffic-nkwxe.mongodb.net/todo?retryWrites=true&w=majority"

	if os.Getenv("DATABASE_URL") != "" {
		connectionString = os.Getenv("DATABASE_URL")
	}

	opts := options.Client()
	opts.ApplyURI(connectionString)
	opts.SetMaxPoolSize(5)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, opts)

	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return ctx, client
}

// Disconnect : Disconnect
/*func Disconnect() {
	fmt.Println("Disconnecting from MongoDB!")
	defer Client.Disconnect(context.Background())
}*/
