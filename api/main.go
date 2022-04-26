package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"context"

	"github.com/NebuDev14/Diseasy-Peasy/api/prisma/db"
	"github.com/gin-gonic/gin"
)

type SymptomQuery struct {
	Symptoms	[]string
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


func getDiseaseByPart(c *gin.Context) {
	part := c.Param("part")
	disease := FindDiseaseByPart(part)
	c.IndentedJSON(http.StatusOK, disease)
}

func getDiseaseByName(c *gin.Context) {
	name := c.Param("name")
	disease := FindDiseaseByName(name)
	c.IndentedJSON(http.StatusOK, disease)
}

func getDiseaseBySymptoms(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil { fmt.Println(err) }
	var data SymptomQuery
	json.Unmarshal(jsonData, &data)
	disease := FindDiseaseBySymptoms(data.Symptoms)

	c.IndentedJSON(http.StatusOK, disease)
}

func matchDisease(c *gin.Context) {
	part := c.Param("part")
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil { fmt.Println(err) }
	var data SymptomQuery
	json.Unmarshal(jsonData, &data)
	diseases := MatchDisease(part, data.Symptoms)

	c.IndentedJSON(http.StatusOK, diseases)
}

func CreateDisease(name string, part string, symptoms []string) error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	newDisease, err := client.Disease.CreateOne(
		db.Disease.Name.Set(name),
		db.Disease.Part.Set(part),
		db.Disease.Symptoms.Set(symptoms),
	).Exec(ctx)

	if err != nil {
		return err
	}

	fmt.Println(newDisease)
	return nil
}

func MatchDisease(part string, symptoms []string) [] db.DiseaseModel{
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	disease, err := client.Disease.FindMany(
		db.Disease.Symptoms.HasSome(symptoms),
		db.Disease.Part.Equals(part),
	).Exec(ctx)

	if err != nil {
		panic(err)
	}

	return disease
}

func FindDiseaseByPart(part string) [] db.DiseaseModel {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	diseases, err := client.Disease.FindMany(
		db.Disease.Part.Equals(part),
	).Exec(ctx)

	if err != nil {
		panic(err)
	}

	return diseases
}

func FindDiseaseByName(name string) *db.DiseaseModel {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()
	disease, err := client.Disease.FindUnique(
		db.Disease.Name.Equals(name),
	).Exec(ctx)

	if err != nil {
		panic(err)
	}

	return disease
}

func FindDiseaseBySymptoms(symptoms []string) [] db.DiseaseModel {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	disease, err := client.Disease.FindMany(
		db.Disease.Symptoms.HasSome(symptoms),
	).Exec(ctx)

	if err != nil {
		panic(err)
	}

	return disease
}
