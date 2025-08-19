package main

import (
	"encoding/json"
	"fmt"
	"library/models"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books []models.Book

func main() {
	loadBooks()
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

	newBook.ID = len(books) + 1
	books = append(books, newBook)
	saveBooks()
	c.JSON(200, newBook)
}

func loadBooks() {
	file, err := os.Open("data/books.json")
	if err != nil {
		fmt.Println("Error file:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&books); err != nil {
		fmt.Println("Error decoding Json: ", err)
	}
}

func saveBooks() {
	file, err := os.Create("data/books.json")
	if err != nil {
		fmt.Println("Error file:", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)

	if err := encoder.Encode(books); err != nil {
		fmt.Println("Error decoding Json: ", err)
	}
}
