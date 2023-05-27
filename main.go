package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "errors"
)
type Book struct {
    ID     string  `json:"id"`  
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Quantity int   `json:"quantity"`
}   

var books = []Book{
    {ID: "1", Title: "Golang pointers", Author: "Mr. Golang", Quantity: 10},
    {ID: "2", Title: "Goroutines", Author: "Mr. Goroutine", Quantity: 20},
    {ID: "3", Title: "Golang routers", Author: "Mr. Router", Quantity: 30},
    {ID: "4", Title: "Golang concurrency", Author: "Mr. Currency", Quantity: 40},}

func GetBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}   

func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range books {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
  
router := gin.Default()
router.GET("/books", GetBooks)
router.GET("/books/:id", GetBookByID)

router.Run("localhost:8080")
}