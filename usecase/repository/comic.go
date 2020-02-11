package repository

import (
	"github.com/alexander-bautista/go-api-2/domain/model"
)

type ComicRepository interface {
	GetOne(id int) (*model.Comic, error)
	GetAll() ([]*model.Comic, error)
}
