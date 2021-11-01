package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	
	"github.com/NebuDev14/Diseasy-Peasy/lib"
	"github.com/gin-gonic/gin"
)

type SymptomQuery struct {
	Symptoms	[]string
}

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
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil { fmt.Println(err) }
	var data SymptomQuery
	json.Unmarshal(jsonData, &data)
	fmt.Println(data.Symptoms[0])
	
	c.IndentedJSON(http.StatusOK, "no")
}

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