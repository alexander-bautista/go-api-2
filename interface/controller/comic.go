package controller

import (
	"github.com/alexander-bautista/go-api-2/domain/model"
	"github.com/alexander-bautista/go-api-2/usecase/interactor"
)

type comicController struct {
	ComicInteractor interactor.ComicInteractor
}

type ComicController interface {
	GetAll() ([]*model.Comic, error)
	GetOne(id int) (*model.Comic, error)
}

func NewComicController(ci interactor.ComicInteractor) ComicController {
	return &comicController{ComicInteractor: ci}
}

func (cc *comicController) GetAll() ([]*model.Comic, error) {
	c, err := cc.ComicInteractor.GetAll()

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cc *comicController) GetOne(id int) (*model.Comic, error) {
	c, err := cc.ComicInteractor.GetOne(id)

	if err != nil {
		return nil, err
	}

	return c, nil

}
