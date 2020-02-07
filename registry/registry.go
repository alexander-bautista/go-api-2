package registry

import (
	"github.com/alexander-bautista/go-api-2/interface/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

type registry struct {
	db *mongo.Client
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *mongo.Client) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewComicController()
}
