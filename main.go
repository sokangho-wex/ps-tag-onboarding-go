package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/handlers"
	"github.com/sokangho-wex/ps-tag-onboarding-go/persistence"
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

	mongoClient := persistence.NewMongoClient(uri)
	db := mongoClient.NewMongoDB()
	defer mongoClient.DisconnectMongoDB()

	// TODO: Find a better way to do dependency injection
	userRepo := persistence.NewUserRepo(db)
	userHandler := handlers.NewUserHandler(userRepo)

	router := gin.Default()
	router.GET("/find/:id", userHandler.FindUser)
	router.POST("/save", userHandler.SaveUser)

	log.Fatal(router.Run(":" + port))
}
