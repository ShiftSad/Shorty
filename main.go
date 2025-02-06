package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strings"
)

type Url struct {
	ID    uint   `gorm:"primaryKey"`
	Short string `gorm:"uniqueIndex"`
	Long  string
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
	short = strings.ToLower(short)

	var url Url
	if err := db.Where("short = ?", short).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if !strings.HasPrefix(url.Long, "https://") && !strings.HasPrefix(url.Long, "http://") {
		url.Long = "https://" + url.Long
	}

	c.Redirect(http.StatusMovedPermanently, url.Long)
}

func shortenURL(c *gin.Context) {
	var request struct {
		URL string `json:"url"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortCode := generateShortCode()
	url := Url{Short: shortCode, Long: request.URL}

	db.Create(&url)
}

// Return a random 6 character string
func generateShortCode() string {
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"
	shortCode := make([]byte, 6)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortCode)
}
