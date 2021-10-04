package main

import (
	"github.com/Fhoust/Go-app/common"
	"github.com/Fhoust/Go-app/database"
	"github.com/Fhoust/Go-app/routes"
)

func main() {

	common.SetupENV()
	database.SetupDB()

	routes.Routes()

}
