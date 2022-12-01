package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/connection"
	_ "github.com/rezairfanwijaya/go-exam-api.git/docs"
	"github.com/rezairfanwijaya/go-exam-api.git/route"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Swagger GO-Exam API Documentation
//	@version		1.0
//	@description	dokumentasi dari endpoint yang dapat digunakan pada aplikasi GO-Exam

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

	// swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start server echo
	if err := e.Start("localhost:9090"); err != nil {
		log.Println(err)
	}
}
