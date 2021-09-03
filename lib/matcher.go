package lib

import (
	"context"
	"fmt"

	"github.com/NebuDev14/Diseasy-Peasy/lib/prisma/db"
)

func MatchDisease(disease string) string {
	if err := run(); err != nil {
        panic(err)
    }
	message := fmt.Sprintf("test")
	fmt.Println("gaming")
	return message
}

func run() error { 
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

	Symptoms := []string {"coughing", "swelling", "pain"}
	
	newDisease, err := client.Disease.CreateOne(
		db.Disease.Name.Set("test1"),
		db.Disease.Part.Set("heart"),
		db.Disease.Symptoms.Set(Symptoms),
	).Exec(ctx)

	if err != nil {
		return err
	}

	fmt.Println(newDisease)
	return nil
}

