package repository

import (
	"context"

	"github.com/alexander-bautista/go-api-2/comic"
	"github.com/alexander-bautista/go-api-2/domain/model"

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

func (m *mongoComicRepository) GetOne(ctx context.Context, id int) (comic *model.Comic, err error) {
	err = m.col.FindOne(ctx, bson.M{"id": id}).Decode(&comic)

	if err != nil {
		return nil, err
	}

	return comic, err
}

func (m *mongoComicRepository) GetAll(ctx context.Context) ([]*model.Comic, error) {

	opts := options.Find()
	//opts.SetLimit(20)

	cursor, _ := m.col.Find(ctx, bson.M{}, opts)

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
