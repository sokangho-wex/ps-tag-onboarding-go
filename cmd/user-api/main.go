package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers/validators"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/persistence"
	"log"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	uri := os.Getenv("MONGO_CONNECTION_STRING")
	if uri == "" {
		uri = "mongodb://root:password@localhost:27017"
	}

	mongoClient := persistence.NewMongoClient(uri)
	db := mongoClient.NewMongoDB()
	defer mongoClient.DisconnectMongoDB()

	userRepo := persistence.NewUserRepo(db)
	userValidator := validators.NewUserValidator(userRepo)
	userHandler := handlers.NewUserHandler(userRepo, userValidator)

	router := gin.Default()
	router.Use(handlers.ErrorHandler())
	router.GET("/find/:id", userHandler.FindUser)
	router.POST("/save", userHandler.SaveUser)

	log.Fatal(router.Run(":" + port))
}