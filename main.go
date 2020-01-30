package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/alexander-bautista/go-api-2/db"
	"github.com/alexander-bautista/go-api-2/models/item"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ctx context.Context
var _client *mongo.Client
var _cancel context.CancelFunc

func main() {
	_ctx, _client = db.Connect()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/items/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
			return
		}

		result := item.GetOne(_client, int(id))

		fmt.Println("item", result)

		if result.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find item with id %s", idParam)})
			return
		}

		comicResponse := getOneMarvelComic()

		//c.JSON(http.StatusOK, result)
		c.JSON(http.StatusOK, comicResponse)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//router.Run()

	/*http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

	_ctx = context.Background()
	_ctx, _cancel = context.WithTimeout(_ctx, time.Second)
	sleepAndTalk(_ctx, 5*time.Second, "hello")

	defer _cancel()*/

	/*_ctx, _client = db.Connect()
	items := item.GetAll(_client)
	fmt.Println(items)

	fmt.Println("get one ", item.GetOne(_client, 1))

	/*newItem := item.Item{ID: 3, Title: "third", IsDone: false}
	fmt.Println("new item id", item.Add(_client, newItem))*/

	//fmt.Println("delete one", item.RemoveOne(_client, 3))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	db.Disconnect(_ctx, _client)
}

func sleepAndTalk(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handler started")

	defer log.Printf("handler ended")

	time.Sleep(5 * time.Second)

	fmt.Fprint(w, "hello")
}

// Item : represents a single database item
type ComicsResponse struct {
	Code   int              `json:"code"`
	Status string           `json:"status,omitempty"`
	Data   ComicDataReponse `json:"data"`
}

type ComicDataReponse struct {
	Total   int           `json:"total"`
	Count   int           `json:"count"`
	Results []ComicResult `json:"results"`
}

type ComicResult struct {
	ID     int    `json:"id"`
	Title  string `json:"title,omitempty"`
	Isbn   string `json:"isbn,omitempty"`
	Format string `json:"format,omitempty"`
}

func getOneMarvelComic() (comicResponse ComicsResponse) {
	publicKey := "ea9048e277273d185472a588dd1f1537"
	hash := "570196dfe16122bcd2ce2819066e93bb"
	ts := 330
	//url := fmt.Sprintf("http://gateway.marvel.com/v1/public/comics?ts=%d&apikey=%s&hash=%s&titleStartsWith=Spider&dateRange=2010-01-01,2013-01-02", ts, publicKey, hash)
	url := fmt.Sprintf("http://gateway.marvel.com/v1/public/comics/40269?ts=%d&apikey=%s&hash=%s", ts, publicKey, hash)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	json.Unmarshal([]byte(body), &comicResponse)

	return
}

func getMarvelData() (comicResponse ComicsResponse) {
	publicKey := "ea9048e277273d185472a588dd1f1537"
	hash := "570196dfe16122bcd2ce2819066e93bb"
	ts := 330
	//url := fmt.Sprintf("http://gateway.marvel.com/v1/public/comics?ts=%d&apikey=%s&hash=%s&titleStartsWith=Spider&dateRange=2010-01-01,2013-01-02", ts, publicKey, hash)
	url := fmt.Sprintf("http://gateway.marvel.com/v1/public/comics/40269?ts=%d&apikey=%s&hash=%s", ts, publicKey, hash)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	log.Println(string(body))

	json.Unmarshal([]byte(body), &comicResponse)

	return
}
