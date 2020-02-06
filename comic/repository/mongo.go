package repository

import (
	"context"

	"github.com/alexander-bautista/go-api-2/comic"
	"github.com/alexander-bautista/go-api-2/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Col = "comics"
)

type mongoComicRepository struct {
	col *mongo.Collection
}

func NewMongoComicRepository(col *mongo.Collection) comic.Repository {
	return &mongoComicRepository{col}
}

func (m *mongoComicRepository) GetOne(ctx context.Context, id int) (comic *models.Comic, err error) {
	err = m.col.FindOne(ctx, bson.M{"id": id}).Decode(&comic)

	if err != nil {
		return nil, err
	}

	return comic, err
}

func (m *mongoComicRepository) GetAll(ctx context.Context) ([]*models.Comic, error) {

	opts := options.Find()
	//opts.SetLimit(20)

	cursor, _ := m.col.Find(ctx, bson.M{}, opts)

	defer cursor.Close(ctx)

	items := make([]*models.Comic, 0)

	for cursor.Next(ctx) {
		oneItem := &models.Comic{}
		err := cursor.Decode(&oneItem)

		if err != nil {
			return nil, err
		}

		items = append(items, oneItem)
	}

	return items, nil
}
