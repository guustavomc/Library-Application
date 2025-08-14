package main

import (
	"fmt"
	"library/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: 1, Name: "LOTR 1", Price: 50.0},
	{ID: 2, Name: "GOT", Price: 33.1},
	{ID: 3, Name: "Star wars", Price: 70.0},
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.GET("/books/:id", getBooksById)

	router.Run()
}

func getBooks(c *gin.Context) {

	c.JSON(200, gin.H{
		"book": books,
	})

}

func getBooksById(c *gin.Context) {
	idParam := c.Param("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	for _, p := range books {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}

	c.JSON(404, gin.H{"message": "Book Not Found"})
}

func postBooks(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}

	books = append(books, newBook)
}
