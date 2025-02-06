package main

import (
	"github.com/gin-gonic/gin"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Short string `json:"short"`
	Long  string `json:"long"`
}

var db *gorm.DB

func main() {
	// Initialize the database
	dsn := os.Getenv("DSN")

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Url{})

	// Initialize the router
	r := gin.Default()

	r.GET(":short", redirectURL)
	r.POST("/shorten", shortenURL)

	r.Run("localhost:8080")
}

func redirectURL(c *gin.Context) {
	short := c.Param("short")

	var url Url
	db.First(&url, "short = ?", short)

	c.Redirect(301, url.Long)
}

func shortenURL(c *gin.Context) {
	var url Url
	err := c.BindJSON(&url)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	db.Create(&url)
	c.JSON(200, gin.H{"short": url.Short})
}
