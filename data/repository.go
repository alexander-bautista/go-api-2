package repository

import (
	"github.com/alexander-bautista/go-api-2/models"
)

type ComicRepository interface {
	//Remove(id int) (Comic, error)
	GetOne(id int) (*models.Comic, error)
	GetAll() ([]*models.Comic, error)
}
