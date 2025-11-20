package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	// No confiar en todos los proxies para evitar el warning
	router.SetTrustedProxies(nil)
	router.GET("/albums", getAlbum)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Escuchar en 0.0.0.0 para que la plataforma detecte el puerto
	addr := ":" + port // Gin usa 0.0.0.0 por defecto cuando se omite host
	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
