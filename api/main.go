package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "gaming")
}

func main() {
	router := gin.Default()
	router.GET("/", test)

	router.Run("localhost:8080")
}