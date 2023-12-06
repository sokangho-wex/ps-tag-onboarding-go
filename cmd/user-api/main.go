package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/onboardingerrors"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/domain/users"
	"github.com/sokangho-wex/ps-tag-onboarding-go/internal/mongo"
	"log"
	"os"
)

const userDB = "user"

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("App unable to start, APP_PORT environment variable is not set")
	}
	uri := os.Getenv("MONGO_CONNECTION_STRING")
	if uri == "" {
		log.Fatal("App unable to start, MONGO_CONNECTION_STRING environment variable is not set")
	}

	mongoClient := mongo.NewClient(uri)
	db := mongoClient.CreateDB(userDB)
	defer mongoClient.DisconnectDB()

	userRepo := mongo.NewUserRepo(db)
	userValidator := users.NewValidator(userRepo)
	userHandler := users.NewHandler(userRepo, userValidator)

	router := gin.Default()
	router.Use(onboardingerrors.ErrorHandler())
	router.GET("/find/:id", userHandler.FindUser)
	router.POST("/save", userHandler.SaveUser)

	log.Fatal(router.Run(":" + port))
}
