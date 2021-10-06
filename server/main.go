package main

import (
	"github.com/Fhoust/Go-app/server/common"
	"github.com/Fhoust/Go-app/server/database"
	"github.com/Fhoust/Go-app/server/routes"
)

func main() {

	common.SetupENV()
	database.SetupDB()

	routes.Routes()

}
