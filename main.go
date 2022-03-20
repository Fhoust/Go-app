package main

import (
	"github.com/zerepl/go-app/common"
	"github.com/zerepl/go-app/database"
	"github.com/zerepl/go-app/routes"
)

func main() {

	common.SetupENV()
	database.SetupDB()

	routes.Routes()

}
