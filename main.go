package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexander-bautista/go-api-2/comics"

	"github.com/gin-gonic/gin"
)

/*var _ctx context.Context
var _client *mongo.Client*/
var _cancel context.CancelFunc

func main() {
	//_ctx, _client = db.Connect()

	router := gin.Default()

	v1 := router.Group("/api")

	comics.ComicsRegister(v1.Group("/comics"))

	test := router.Group("/api/ping")

	test.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
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

	//db.Disconnect()
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
