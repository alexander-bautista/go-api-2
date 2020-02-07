package comic

import (
	"context"

	"github.com/alexander-bautista/go-api-2/domain/model"
)

type Usecase interface {
	GetOne(ctx context.Context, id int) (*model.Comic, error)
	GetAll(ctx context.Context) ([]*model.Comic, error)
	EstimatedTaxes(ctx context.Context, id int) (float32, error)
}
