package main

import (
	"github.com/gin-gonic/gin"
	errorhandler "github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/onboardingerrors/httphandler"
	userhandler "github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/users/httphandler"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/users/validator"
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
	userValidator := validator.NewUserValidator(userRepo)
	userHandler := userhandler.NewUserHandler(userRepo, userValidator)

	router := gin.Default()
	router.Use(errorhandler.ErrorHandler())
	router.GET("/find/:id", userHandler.FindUser)
	router.POST("/save", userHandler.SaveUser)

	log.Fatal(router.Run(":" + port))
}
