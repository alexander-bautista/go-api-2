package repository

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/alexander-bautista/go-api-2/domain/model"
	"github.com/alexander-bautista/go-api-2/usecase/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoComicRepository(mongoURL, mongoDB string, mongoTimeout int) (repository.ComicRepository, error) {
	repo := &mongoRepository{
		database: mongoDB,
		timeout:  time.Duration(mongoTimeout) * time.Second,
	}

	client, err := newMongoClient(mongoURL, mongoTimeout)

	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}

	repo.client = client
	return repo, nil

}

func (r *mongoRepository) GetOne(id int) (comic *model.Comic, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()

	collection := r.client.Database("todo").Collection("comics")

	err = collection.FindOne(ctx, bson.M{"id": id}).Decode(&comic)

	if err != nil {
		return nil, errors.Wrap(err, "repository.Comic.GetOne")
	}

	return comic, err
}

func (r *mongoRepository) GetAll() ([]*model.Comic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()

	opts := options.Find()
	//opts.SetLimit(20)

	collection := r.client.Database("todo").Collection("comics")

	cursor, _ := collection.Find(ctx, bson.M{}, opts)

	defer cursor.Close(ctx)

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
