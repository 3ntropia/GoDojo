package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albumsDe", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
