package presenter

import (
	"github.com/alexander-bautista/go-api-2/domain/model"
)

type ComicPresenter interface {
	GetAllComics(c []*model.Comic) []*model.Comic
	GetOneComic(c *model.Comic) *model.Comic
}
