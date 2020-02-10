package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexander-bautista/go-api-2/infrastructure/datastore"
	ro "github.com/alexander-bautista/go-api-2/infrastructure/router"
	"github.com/alexander-bautista/go-api-2/registry"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var _cancel context.CancelFunc

func main() {

	router := gin.Default()

	v1 := router.Group("/api")

	ctx, col := datastore.Connect()

	//defer Disconnect(ctx, )

	r := registry.NewRegistry(col)

	ro.NewRouter(v1, r.NewAppController())

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
		log.Fatal("Error while shutdown server:", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	//Disconnect(ctx, col)
	log.Println("Server exiting")
}

// Disconnect : Disconnect
func Disconnect(ctx context.Context, col *mongo.Collection) {
	fmt.Println("Disconnecting from MongoDB!")
	defer col.Database().Client().Disconnect(ctx)
}
