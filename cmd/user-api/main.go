package main

import (
	"github.com/gin-gonic/gin"
	handlers2 "github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers/validators"
	persistence2 "github.com/sokangho-wex/ps-tag-onboarding-go/internal/persistence"
	"log"
	"os"
)

func main() {
	// TODO: Move "getting config values" code to a separate method or file
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	uri := os.Getenv("MONGO_CONNECTION_STRING")
	if uri == "" {
		uri = "mongodb://root:password@localhost:27017"
	}

	mongoClient := persistence2.NewMongoClient(uri)
	db := mongoClient.NewMongoDB()
	defer mongoClient.DisconnectMongoDB()

	// TODO: Find a better way to do dependency injection
	userRepo := persistence2.NewUserRepo(db)
	userValidator := validators.NewUserValidator(userRepo)
	userHandler := handlers2.NewUserHandler(userRepo, userValidator)

	router := gin.Default()
	router.Use(handlers2.ErrorHandler())
	router.GET("/find/:id", userHandler.FindUser)
	router.POST("/save", userHandler.SaveUser)

	log.Fatal(router.Run(":" + port))
}