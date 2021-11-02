package lib

import (
	"context"
	"fmt"
	"github.com/NebuDev14/Diseasy-Peasy/lib/prisma/db"
)

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

func MatchDisease(part string, symptoms []string) {
	diseasesBySymps := FindDiseaseBySymptoms(symptoms)
	var disease db.DiseaseModel[]
	for i, s := range db.DiseaseModel {

	}
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
