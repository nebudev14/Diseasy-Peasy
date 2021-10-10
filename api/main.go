package main

import (
	"fmt"
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

func getSymptom(c *gin.Context) {
	name := c.Param("name")
	symptom := lib.FindSymptom(name)
	c.IndentedJSON(http.StatusOK, symptom)
}

func main() {

	fmt.Println(lib.FindDiseaseByName("test2"))
	// fmt.Println(lib.FindDiseaseBySymptoms("test1"))
	// lib.CreateSymptom("vomit", "test1")
	router := gin.Default()
	router.GET("/parts/:part", getDiseaseByPart)
	router.GET("/name/:name", getDiseaseByName)
	router.GET("/symptom/:name", getSymptom)
	router.Run("localhost:8080")
}