package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/handlers"
	"log"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/find/:id", handlers.FindUser)
	router.POST("/save", handlers.SaveUser)

	log.Fatal(router.Run(":" + port))
}
