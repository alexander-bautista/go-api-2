package registry

import (
	"log"
	"os"
	"strconv"

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

func (r *registry) NewComicInteractor() interactor.ComicInteractor {
	return interactor.NewComicInteractor(r.NewComicRepository(), r.NewComicPresenter())
}

func (r *registry) NewComicRepository() cr.ComicRepository {
	return chooseRepo()
}

func (r *registry) NewComicPresenter() cp.ComicPresenter {
	return p.NewComicPresenter()
}

func chooseRepo() cr.ComicRepository {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb+srv://todo_user:todo2020@traffic-nkwxe.mongodb.net/todo?retryWrites=true&w=majority"
	}
	mongodb := os.Getenv("MONGO_DB")
	if mongodb == "" {
		mongodb = "todo"
	}

	mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))

	if mongoTimeout == 0 {
		mongoTimeout = 10
	}

	repo, err := rp.NewMongoComicRepository(mongoURL, mongodb, mongoTimeout)

	if err != nil {
		log.Fatal(err)
	}

	return repo

}
