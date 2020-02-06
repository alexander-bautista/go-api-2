package comic

import (
	"context"
	"github.com/alexander-bautista/go-api-2/models"
)

type Usecase interface {
	GetOne(ctx context.Context, id int) (*models.Comic, error)
	GetAll(ctx context.Context) ([]*models.Comic, error)
	EstimatedTaxes(ctx context.Context, id int) (float32, error)
}
