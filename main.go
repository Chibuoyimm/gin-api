package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

var albums = []album{
	{ID: "1", Title: "To pimp a butterfly", Artist: "Kendrick Lamar", Price: 100.5},
	{ID: "2", Title: "good kid, m.A.A.d city", Artist: "Kendrick Lamar", Price: 200},
	{ID: "3", Title: "Damn", Artist: "Kendrick Lamar", Price: 80},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, album := range(albums) {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for index, album := range(albums) {
		if album.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}
	for index, album := range(albums) {
		if album.ID == id {
			albums[index] = updatedAlbum
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)

	router.Run()
}
