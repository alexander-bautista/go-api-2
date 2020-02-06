package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_comicHttpDeliver "github.com/alexander-bautista/go-api-2/comic/delivery/http"
	_comicRepo "github.com/alexander-bautista/go-api-2/comic/repository"
	_comicUsecase "github.com/alexander-bautista/go-api-2/comic/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gin-gonic/gin"
)

var _cancel context.CancelFunc

func main() {
	router := gin.Default()
	v1 := router.Group("/api")

	ctx, col := Connect()

	comicRepo := _comicRepo.NewMongoComicRepository(col)
	timeoutContext := time.Duration(5 * time.Second)

	cu := _comicUsecase.NewComicUsecase(comicRepo, timeoutContext)

	_comicHttpDeliver.NewComicHandler(v1.Group("/comics"), cu)

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

	Disconnect(ctx, col)
	log.Println("Server exiting")
}

func Connect() (context.Context, *mongo.Collection) {
	connectionString := "mongodb+srv://todo_user:todo2020@traffic-nkwxe.mongodb.net/todo?retryWrites=true&w=majority"

	if os.Getenv("DATABASE_URL") != "" {
		connectionString = os.Getenv("DATABASE_URL")
	}

	opts := options.Client()
	opts.ApplyURI(connectionString)
	opts.SetMaxPoolSize(5)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, opts)

	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return ctx, client.Database("todo").Collection("comics")
}

// Disconnect : Disconnect
func Disconnect(ctx context.Context, col *mongo.Collection) {
	fmt.Println("Disconnecting from MongoDB!")
	defer col.Database().Client().Disconnect(ctx)
}
