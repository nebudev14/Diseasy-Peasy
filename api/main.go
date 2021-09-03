package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/NebuDev14/Diseasy-Peasy/lib"

)

func test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "gaming")
}

func main() {

	// sypmtoms := []string {"pain", "bloody nose"}

	fmt.Println(lib.FindDiseaseByPart("heart")[0])
	

	// router := gin.Default()
	// router.GET("/", test)
	// router.Run("localhost:8080")
}