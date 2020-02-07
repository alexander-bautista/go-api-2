package repository

type ComicRepository interface {
	//Remove(id int) (Comic, error)
	GetOne(id int) (*models.Comic, error)
	GetAll() ([]*models.Comic, error)
}
