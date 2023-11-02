package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers/onboardingerrors"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers/users"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/handlers/validators"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/persistence"
	"log"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("App unable to start, APP_PORT environment variable is not set")
	}
	uri := os.Getenv("MONGO_CONNECTION_STRING")
	if uri == "" {
		log.Fatal("App unable to start, MONGO_CONNECTION_STRING environment variable is not set")
	}

	mongoClient := persistence.NewMongoClient(uri)
	db := mongoClient.NewMongoUserDB()
	defer mongoClient.DisconnectMongoDB()

	userRepo := persistence.NewUserRepo(db)
	userValidator := validators.NewUserValidator(userRepo)
	userHandler := users.NewUserHandler(userRepo, userValidator)

	router := gin.Default()
	router.Use(onboardingerrors.ErrorHandler())
	router.GET("/find/:id", userHandler.FindUser)
	router.POST("/save", userHandler.SaveUser)

	log.Fatal(router.Run(":" + port))
}
