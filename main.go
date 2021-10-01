package main

import (
	"log"

	"github.com/Fhoust/Go-app/routes"
	common "github.com/Fhoust/Go-app/handlers"
)

func main() {

	common.SetupENV()
	common.SetupDB()

	routes.Routes()


	log.Println("Initializing the app")

}
