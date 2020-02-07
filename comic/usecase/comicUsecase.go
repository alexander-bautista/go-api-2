package usecase

import (
	"context"
	"time"

	"github.com/alexander-bautista/go-api-2/comic"
	"github.com/alexander-bautista/go-api-2/domain/model"
)

type comicUsecase struct {
	repo           comic.Repository
	contextTimeout time.Duration
}

func NewComicUsecase(c comic.Repository, timeout time.Duration) comic.Usecase {
	return &comicUsecase{
		repo:           c,
		contextTimeout: timeout,
	}
}

func (c *comicUsecase) GetOne(ctx context.Context, id int) (*model.Comic, error) {
	con, cancel := context.WithTimeout(ctx, c.contextTimeout)

	defer cancel()

	comic, err := c.repo.GetOne(con, id)

	if err != nil {
		return nil, err
	}

	return comic, nil
}

func (c *comicUsecase) GetAll(ctx context.Context) ([]*model.Comic, error) {
	con, cancel := context.WithTimeout(ctx, c.contextTimeout)

	defer cancel()

	comics, err := c.repo.GetAll(con)

	if err != nil {
		return nil, err
	}

	return comics, nil
}

func (c *comicUsecase) EstimatedTaxes(ctx context.Context, id int) (float32, error) {
	con, cancel := context.WithTimeout(ctx, c.contextTimeout)

	defer cancel()

	comic, err := c.repo.GetOne(con, id)

	if err != nil {
		return 0, err
	}

	return comic.EstimatedTaxes(), nil

}
