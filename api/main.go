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

// func getSymptom(c *gin.Context) {
// 	name := c.Param("name")
// 	symptom := lib.FindSymptom(name)
// 	c.IndentedJSON(http.StatusOK, symptom)
// }

func main() {
	
	fmt.Println(lib.FindDiseaseByName("test2"))
	symps := []string{"memory loss", "pain"};
	lib.CreateDisease("Alzheimers", "Brain", symps)
	fmt.Println(lib.FindDiseaseBySymptoms(symps))
	router := gin.Default()
	router.GET("/parts/:part", getDiseaseByPart)
	router.GET("/name/:name", getDiseaseByName)
	// router.GET("/symptom/:name", getSymptom)
	router.Run("localhost:8080")
}