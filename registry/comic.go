package registry

import (
	"github.com/alexander-bautista/go-api-2/interface/controller"
	p "github.com/alexander-bautista/go-api-2/interface/presenter"
	rp "github.com/alexander-bautista/go-api-2/interface/repository"
	"github.com/alexander-bautista/go-api-2/usecase/interactor"
	cp "github.com/alexander-bautista/go-api-2/usecase/presenter"
	cr "github.com/alexander-bautista/go-api-2/usecase/repository"
)

func (r *registry) NewComicController() controller.ComicController {
	return controller.NewComicController(r.NewComicInteractor())
}


func (r * registry) NewComicInteractor() interactor.ComicInteractor {
	return interactor.NewComicInteractor(r.NewComicRepository(), r.NewComicPresenter())
}

func (r *registry)  NewComicRepository() cr.ComicRepository {
	return rp.NewMongoComicRepository(r.db)
}

func (r *registry) NewComicPresenter() cp.ComicPresenter {
	return  p.NewComicPresenter()
}