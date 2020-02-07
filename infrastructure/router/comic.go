package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexander-bautista/go-api-2/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.RouterGroup, c controller.AppController) {
	//router.GET("/", retrieveAll)
	router.GET("/:id", retrieveOne)
}

func retrieveOne(g *gin.Context, c controller.AppController) {
	idParam := g.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	comic, err := c.GetOne(g, int(id))

	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find item with id %s", idParam)})
	}

	g.JSON(http.StatusOK, comic)
}

/*
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

	c.JSON(http.StatusOK, comics)
}
*/
