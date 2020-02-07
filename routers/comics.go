package routers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexander-bautista/go-api-2/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	serv services.ComicService
}

func (h Handler) Register(router *gin.RouterGroup) {
	router.GET("/", h.retrieveAll)
	router.GET("/:id", h.retrieveOne)
}

func (h Handler) retrieveOne(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	comic, err := h.serv.GetOne(int(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find item with id %s", idParam)})
	}

<<<<<<< HEAD:comics/routers.go
	comicResponse := findComic(idParam)
=======
	//comicResponse := GetOneComicById(idParam)
>>>>>>> e57e652facf8869d7604002992d09f0dab6cd42e:routers/comics.go

	c.JSON(http.StatusOK, comic)
}

func (h Handler) retrieveAll(c *gin.Context) {
	// Get query  parameters.
	dateRange := c.Query("dateRange")
	titleStartsWith := c.Query("titleStartsWith")

	if len(dateRange) == 0 && len(titleStartsWith) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Must provide a filter dateRange or titleStartsWith"})
	}

	comics, err := h.serv.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Get Marvel data (always) with filters
<<<<<<< HEAD:comics/routers.go
	comicsResponse := findComics(titleStartsWith, dateRange)
=======
	//comicsResponse := GetComics(titleStartsWith, dateRange)
>>>>>>> e57e652facf8869d7604002992d09f0dab6cd42e:routers/comics.go

	/*if comicsResponse.Code != http.StatusOK {
		log.Fatal("Marvel API response code", comicsResponse.Code)
	}*/

	//comics := comicsResponse.Data.Results

	//items, err := findAndUpdateMany(comics)

	/*if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}*/

	c.JSON(http.StatusOK, comics)
}
