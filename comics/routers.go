package comics

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"strconv"
)

func ComicsRegister(router *gin.RouterGroup) {
	router.GET("/", RetrieveAll)
	router.GET("/:id", RetrieveOne)
}

func RetrieveOne(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	result := GetOne(int(id))

	if result.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find item with id %s", idParam)})
		return
	}

	comicResponse := GetOneComicById(idParam)

	c.JSON(http.StatusOK, comicResponse)
}

func RetrieveAll(c *gin.Context) {
	// Get query  parameters.
	dateRange := c.Query("dateRange")
	titleStartsWith := c.Query("titleStartsWith")

	if len(dateRange) == 0 && len(titleStartsWith) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Must provide a filter dateRange or titleStartsWith"})
	}

	// Get Marvel data (always) with filters
	comicsResponse := GetComics(titleStartsWith, dateRange)

	if comicsResponse.Code != http.StatusOK {
		log.Fatal("Marvel API response code", comicsResponse.Code)
	}

	comics := comicsResponse.Data.Results
	// Send a findOneAndUpdate with upsert = true and returnNewDocument = true

	var items = make([]interface{}, len(comics))

	for i, comic := range comics {
		doc, err := FindOneAndUpdate(comic)

		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}

		items[i] = doc
	}

	/*items := make([]interface{}, len(comics))
	  for i := range comics {
	  	comics[i].Quantity = rand.Intn(1000)
	  	items[i] = comics[i]
	  }
	  comic.AddMany(_client, items)*/

	c.JSON(http.StatusOK, items)
}
