package interactor

import (
	"github.com/alexander-bautista/go-api-2/domain/model"
	"github.com/alexander-bautista/go-api-2/usecase/presenter"
	"github.com/alexander-bautista/go-api-2/usecase/repository"
)

type comicInteractor struct {
	Repository repository.ComicRepository
	Presenter  presenter.ComicPresenter
}

type ComicInteractor interface {
	GetAll() ([]*model.Comic, error)
	GetOne(id int) (*model.Comic, error)
}

func NewComicInteractor(r repository.ComicRepository, p presenter.ComicPresenter) ComicInteractor {
	return &comicInteractor{r, p}
}

func (ci *comicInteractor) GetAll() ([]*model.Comic, error) {
	c, err := ci.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return ci.Presenter.GetAllComics(c), nil
}

func (ci *comicInteractor) GetOne(id int) (*model.Comic, error) {
	c, err := ci.Repository.GetOne(id)
	if err != nil {
		return nil, err
	}

	return ci.Presenter.GetOneComic(c), nil
}
