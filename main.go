package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Year: 1957},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Year: 1957},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1954},
}

func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbum)
	router.Run("localhost:8080")
}
