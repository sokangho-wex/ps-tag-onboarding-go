package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sokangho-wex/ps-tag-onboarding-go/handlers"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/find/:id", handlers.FindUser)
	router.POST("/save", handlers.SaveUser)

	log.Fatal(router.Run(":8080"))
}
