package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"test/config"
	"test/internal/controller"
	"test/internal/repository"
	"test/internal/service"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}

	defer mongoClient.Disconnect(context.TODO())

	r := gin.Default()

	clickRepo := repository.NewClickRepository(mongoClient.Database(cfg.MongoDB))
	clickService := service.NewClickService(clickRepo)
	clickController := controller.NewClickController(clickService)

	r.GET("/counter/:id", clickController.Update)
	r.POST("/stats/:id", clickController.GetStats)
	r.POST("/save", clickController.Save)
	slog.Info("Strat Server")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
