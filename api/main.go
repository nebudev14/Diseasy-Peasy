package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/NebuDev14/Diseasy-Peasy/lib"

)

func getDiseaseByPart(c *gin.Context) {
	part := c.Param("part")
	disease := lib.FindDiseaseByPart(part)
	c.IndentedJSON(http.StatusOK, disease)
}

func getDiseaseByName(c *gin.Context) {
	name := c.Param("name")
	disease := lib.FindDiseaseByName(name)
	c.IndentedJSON(http.StatusOK, disease)
}

func main() {

	// fmt.Println(lib.FindDiseaseByPart("heart")[0].Symptoms[0])
	// fmt.Println(lib.FindDiseaseByName("test1").Part)
	
	router := gin.Default()
	router.GET("/parts/:part", getDiseaseByPart)
	router.GET("/name/:name", getDiseaseByName)
	router.Run("localhost:8080")
}