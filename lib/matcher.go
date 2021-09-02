package lib

import (
	"fmt"
	"github.com/NebuDev14/Diseasy-Peasy/lib/prisma/db"

)

func MatchDisease(disease string) string {

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return ""
	}

	message := fmt.Sprintf("test")
	fmt.Println("gaming")
	return message
}