package main

import (
	"log"

	"github.com/rezairfanwijaya/go-exam-api.git/connection"
)

func main() {
	DBConnecton, err := connection.DB(".env")
	if err != nil {
		log.Println(err)
	}

	log.Println(DBConnecton)
}
