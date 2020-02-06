package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexander-bautista/go-api-2/comic"
	"github.com/gin-gonic/gin"
)

type ComicHandler struct {
	CUsecase comic.Usecase
}

func NewComicHandler(router *gin.RouterGroup, cu comic.Usecase) {
	handler := &ComicHandler{
		CUsecase: cu,
	}

	router.GET("/", handler.retieveAll)
	router.GET("/:id", handler.retrieveOne)
	router.GET("/:id/estimatedTaxes", handler.estimatedTaxes)
}
func (h ComicHandler) retrieveOne(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	comic, err := h.CUsecase.GetOne(c, int(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find item with id %s", idParam)})
	}

	c.JSON(http.StatusOK, comic)
}

func (h ComicHandler) retieveAll(c *gin.Context) {
	// Get query  parameters.
	dateRange := c.Query("dateRange")
	titleStartsWith := c.Query("titleStartsWith")

	if len(dateRange) == 0 && len(titleStartsWith) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Must provide a filter dateRange or titleStartsWith"})
		return
	}

	comics, err := h.CUsecase.GetAll(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comics)
}

func (h ComicHandler) estimatedTaxes(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	taxes, err := h.CUsecase.EstimatedTaxes(c, int(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find item with id %s", idParam)})
	}

	c.JSON(http.StatusOK, taxes)
}
