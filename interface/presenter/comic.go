package presenter

import "github.com/alexander-bautista/go-api-2/domain/model"

type comicPresenter struct {
}

type ComicPresenter interface {
	ResponseUsers(us []*model.Comic) []*model.Comic
}

func NewComicPresenter() ComicPresenter {
	return &comicPresenter{}
}

func (cp *comicPresenter) ResponseComics(com []*model.Comic) []*model.Comic {
	for _, c := range com {
		c.Title = "Mr." + c.Title
	}
	return com
}
