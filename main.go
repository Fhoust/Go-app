package main

import (
	"log"

	"github.com/Fhoust/Go-app/routes"
	"github.com/Fhoust/Go-app/common"
	"github.com/Fhoust/Go-app/database"
)

func main() {

	common.SetupENV()
	database.SetupDB()

	routes.Routes()


	log.Println("Initializing the app")

}
