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

func getDiseaseBySymptoms(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil { fmt.Println(err) }
	var data SymptomQuery
	json.Unmarshal(jsonData, &data)
	disease := lib.FindDiseaseBySymptoms(data.Symptoms)

	c.IndentedJSON(http.StatusOK, disease)
}

func matchDisease(c *gin.Context) {
	part := c.Param("part")
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil { fmt.Println(err) }
	var data SymptomQuery
	json.Unmarshal(jsonData, &data)
	diseases := lib.MatchDisease(part, data.Symptoms)

	c.IndentedJSON(http.StatusOK, diseases)
}

func main() {
	// symps := []string{"pain", "weak bones"};
	// lib.CreateDisease("osteoporosis", "bones", symps)

	router := gin.Default()
	router.GET("/parts/:part", getDiseaseByPart)
	router.GET("/name/:name", getDiseaseByName)
	router.GET("/symptoms", getDiseaseBySymptoms)
	router.GET("/match/:part", matchDisease)
	router.Run("localhost:8080")
}
