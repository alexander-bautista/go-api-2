package comic

import (
	"context"

	"github.com/alexander-bautista/go-api-2/domain/model"
)

type Repository interface {
	GetOne(ctx context.Context, id int) (*model.Comic, error)
	GetAll(ctx context.Context) ([]*model.Comic, error)
}
