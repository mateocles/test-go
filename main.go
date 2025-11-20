package main

import (
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var (
	albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Year: 1957},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Year: 1957},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1954},
	}
	albumsMu sync.RWMutex
	nextID   = 4
)

// LISTAR todos los álbumes
func getAlbums(c *gin.Context) {
	albumsMu.RLock()
	defer albumsMu.RUnlock()
	c.JSON(http.StatusOK, albums)
}

// OBTENER un álbum por ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	albumsMu.RLock()
	defer albumsMu.RUnlock()
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "album no encontrado"})
}

// CREAR un nuevo álbum
func createAlbum(c *gin.Context) {
	var input struct {
		Title  string `json:"title"`
		Artist string `json:"artist"`
		Year   int    `json:"year"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	if input.Title == "" || input.Artist == "" || input.Year == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title, artist y year son requeridos"})
		return
	}
	albumsMu.Lock()
	id := strconv.Itoa(nextID)
	nextID++
	newAlbum := album{ID: id, Title: input.Title, Artist: input.Artist, Year: input.Year}
	albums = append(albums, newAlbum)
	albumsMu.Unlock()
	c.JSON(http.StatusCreated, newAlbum)
}

// ACTUALIZAR un álbum existente completo (PUT)
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Title  string `json:"title"`
		Artist string `json:"artist"`
		Year   int    `json:"year"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	if input.Title == "" || input.Artist == "" || input.Year == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title, artist y year son requeridos"})
		return
	}
	albumsMu.Lock()
	defer albumsMu.Unlock()
	for i, a := range albums {
		if a.ID == id {
			albums[i] = album{ID: id, Title: input.Title, Artist: input.Artist, Year: input.Year}
			c.JSON(http.StatusOK, albums[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "album no encontrado"})
}

// ELIMINAR un álbum por ID
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	albumsMu.Lock()
	defer albumsMu.Unlock()
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "album no encontrado"})
}

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	// No confiar en todos los proxies para evitar el warning
	router.SetTrustedProxies(nil)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", createAlbum)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

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
