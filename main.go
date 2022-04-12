package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zerepl/go-app/internal/controllers"
	repositories "github.com/zerepl/go-app/internal/data/users"
	"github.com/zerepl/go-app/internal/domain/common"
	"github.com/zerepl/go-app/internal/domain/services/user"
	"github.com/zerepl/go-app/internal/driver"
)

func main() {
	common.SetupENV()

	db := driver.SetupDB()
	defer db.Close()

	router := gin.Default()
	userRepository := repositories.NewUserRepository(db)
	userService := user.NewUserService(userRepository)

	userController := controllers.NewUserController(userService)
	userController.UserRoutes(router)

	router.Run(common.GetPort())
}
