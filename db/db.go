package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DbName :  Database name
const DbName = "todo"

// Connect : Connect
func Connect() (context.Context, *mongo.Client) {

	connectionString := "mongodb+srv://todo_user:todo2020@traffic-nkwxe.mongodb.net/todo?retryWrites=true&w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return ctx, client
}

// Disconnect : Disconnect
func Disconnect(ctx context.Context, client *mongo.Client) {
	fmt.Println("Disconnecting from MongoDB!")
	defer client.Disconnect(ctx)
}
