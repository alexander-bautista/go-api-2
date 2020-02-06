package services

import (
	"fmt"

	data "github.com/alexander-bautista/go-api-2/data"
	"github.com/alexander-bautista/go-api-2/models"
)

type ComicService struct {
	repo data.ComicRepository
}

func (c ComicService) IsDuplicated(id int) error {
	comic, err := c.repo.GetOne(id)

	// comic != nil
	if comic.Id == id {
		return fmt.Errorf("A comic with %d already exist", id)
	}

	if err != nil {
		return err
	}

	return nil

}

func (c *ComicService) GetOne(id int) (*models.Comic, error) {
	return c.repo.GetOne(id)
}

func (c ComicService) GetAll() ([]*models.Comic, error) {
	return c.GetAll()
}

func (c ComicService) EstimatedTaxes(id int) float32 {
	comic, err := c.repo.GetOne(id)

	if err != nil {
		fmt.Println("err", err)
	}

	return comic.EstimatedTaxes()

}
