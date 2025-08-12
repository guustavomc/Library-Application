package main

import (
	"library/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.Run()
}

func getBooks(c *gin.Context) {
	var books = []models.Book{
		{ID: 1, Name: "LOTR 1", Price: 50.0},
		{ID: 2, Name: "GOT", Price: 33.1},
		{ID: 3, Name: "Star wars", Price: 70.0},
	}
	c.JSON(200, gin.H{
		"book": books,
	})

}
