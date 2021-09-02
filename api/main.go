package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/NebuDev14/Diseasy-Peasy/lib"
)

func test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "gaming")
}

func main() {

	lib.MatchDisease("hello")

	router := gin.Default()
	router.GET("/", test)
	router.Run("localhost:8080")
}