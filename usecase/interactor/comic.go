package interactor

import (
	"context"

	"github.com/alexander-bautista/go-api-2/domain/model"
	"github.com/alexander-bautista/go-api-2/usecase/presenter"
	"github.com/alexander-bautista/go-api-2/usecase/repository"
)

type comicInteractor struct {
	Repository repository.ComicRepository
	Presenter  presenter.ComicPresenter
}

type ComicInteractor interface {
	GetAll(ctx context.Context) ([]*model.Comic, error)
	GetOne(ctx context.Context, id int) (*model.Comic, error)
}

func NewComicInteractor(r repository.ComicRepository, p presenter.ComicPresenter) ComicInteractor {
	return &comicInteractor{r, p}
}

func (ci *comicInteractor) GetAll(ctx context.Context) ([]*model.Comic, error) {
	c, err := ci.Repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return ci.Presenter.GetAllComics(c), nil
}

func (ci *comicInteractor) GetOne(ctx context.Context, id int) (*model.Comic, error) {
	c, err := ci.Repository.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return ci.Presenter.GetOneComic(c), nil
}
