package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect : Connect
func Connect() (context.Context, *mongo.Client) {

	connectionString := "mongodb+srv://todo_user:todo2020@traffic-nkwxe.mongodb.net/todo?retryWrites=true&w=majority"

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return ctx, client
}

// Disconnect : Disconnect
func Disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}
