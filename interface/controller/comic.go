package controller

import (
	"context"

	"github.com/alexander-bautista/go-api-2/domain/model"
	"github.com/alexander-bautista/go-api-2/usecase/interactor"
)

type comicController struct {
	ComicInteractor interactor.ComicInteractor
}

type ComicController interface {
	GetAll(ctx context.Context) ([]*model.Comic, error)
	GetOne(ctx context.Context, id int) (*model.Comic, error)
}

func NewComicController(ci interactor.ComicInteractor) ComicController {
	return &comicController{ComicInteractor: ci}
}

func (cc *comicController) GetAll(ctx context.Context) ([]*model.Comic, error) {
	c, err := cc.ComicInteractor.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cc *comicController) GetOne(ctx context.Context, id int) (*model.Comic, error) {
	c, err := cc.ComicInteractor.GetOne(ctx, id)

	if err != nil {
		return nil, err
	}

	return c, nil

}
