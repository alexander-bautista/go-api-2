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

var Ctx context.Context
var Client *mongo.Client
var Cancel context.CancelFunc

// Connect : Connect
func Connect() (context.Context, *mongo.Client) {

	connectionString := "mongodb+srv://todo_user:todo2020@traffic-nkwxe.mongodb.net/todo?retryWrites=true&w=majority"

	Ctx, Cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Client, err := mongo.Connect(Ctx, options.Client().ApplyURI(connectionString))

	defer Cancel()

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(Ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return Ctx, Client
}

// Disconnect : Disconnect
func Disconnect() {
	fmt.Println("Disconnecting from MongoDB!")
	defer Client.Disconnect(Ctx)
}
