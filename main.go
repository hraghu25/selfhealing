package main

import (
	"selfhealing/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize default router
	r := gin.Default()

	// Enabled Logger
	r.Use(gin.Logger())

	// HTTP GET request route
	r.GET("/", handlers.DefaultResponse)

	// Run webserver
	r.Run() // ---> Default to run on 8080
}
