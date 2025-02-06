package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET(":short", redirectURL)
	r.POST("/shorten", shortenURL)

	r.Run("localhost:8080")
}

func redirectURL(c *gin.Context) {

}

func shortenURL(c *gin.Context) {

}
