package lib

import (
	"context"
	"fmt"
	
	"github.com/NebuDev14/Diseasy-Peasy/lib/prisma/db"
)

func MatchDisease(disease string) string {
	message := fmt.Sprintf("test")
	fmt.Println("gaming")
	return message
}

func CreateDisease(name string, part string) error { 
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
	).Exec(ctx)

	if err != nil {
		return err
	}

	fmt.Println(newDisease)
	return nil
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

// func FindDiseaseBySymptoms(symptoms string) [] db.DiseaseModel {
// 	client := db.NewClient()
// 	if err := client.Prisma.Connect(); err != nil {
// 		panic(err)
// 	}

// 	defer func() {
// 		if err := client.Prisma.Disconnect(); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	ctx := context.Background()

// 	disease, err := client.Disease.FindMany(
// 		db.Disease.Symptoms.Contains("pain"),
// 	).Exec(ctx)

	
// 	if err != nil {
// 		panic(err)
// 	}

// 	return disease
// }