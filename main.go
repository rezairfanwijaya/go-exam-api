package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/connection"
	"github.com/rezairfanwijaya/go-exam-api.git/route"
)

func main() {
	// create connection to database
	dbConnection, err := connection.DB(".env")
	if err != nil {
		log.Println(err)
	}

	// define echo
	e := echo.New()

	// define route
	route.NewRoute(e, dbConnection)

	// start server echo
	if err := e.Start("localhost:9090"); err != nil {
		log.Println(err)
	}
}
