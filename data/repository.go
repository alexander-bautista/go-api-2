package repository

import "github.com/alexander-bautista/go-api-2/domain/model"

type ComicRepository interface {
	//Remove(id int) (Comic, error)
	GetOne(id int) (*model.Comic, error)
	GetAll() ([]*model.Comic, error)
}
