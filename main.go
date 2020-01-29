package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/alexander-bautista/go-api-2/db"
	"github.com/alexander-bautista/go-api-2/models/item"
)

var _ctx context.Context
var _client *mongo.Client

func main() {
	_ctx, _client = db.Connect()
	items := item.GetAll(_client)
	fmt.Println(items)

	fmt.Println("get one ", item.GetOne(_client, 1))

	db.Disconnect(_ctx, _client)
}
