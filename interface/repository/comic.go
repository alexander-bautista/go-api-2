package repository

import (
	"context"

	"github.com/alexander-bautista/go-api-2/domain/model"
	"github.com/alexander-bautista/go-api-2/usecase/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Col = "comics"
)

type comicRepository struct {
	col *mongo.Collection
}

func NewMongoComicRepository(col *mongo.Collection) repository.ComicRepository {
	return &comicRepository{col}
}

func (m *comicRepository) GetOne(ctx context.Context, id int) (comic *model.Comic, err error) {
	err = m.col.FindOne(ctx, bson.M{"id": id}).Decode(&comic)

	if err != nil {
		return nil, err
	}

	return comic, err
}

func (m *comicRepository) GetAll(ctx context.Context) ([]*model.Comic, error) {

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
