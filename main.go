package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	setupServer()
}

func setupServer() {
	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
