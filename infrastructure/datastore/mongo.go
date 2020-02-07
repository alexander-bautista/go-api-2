package datastore

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DbName :  Database name
const DbName = "todo"

func Connect() (context.Context, *mongo.Client) {
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
