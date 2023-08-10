package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// The below function will return the default response to the HTTP request on /

func DefaultResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, basic setup is up!"})
}
