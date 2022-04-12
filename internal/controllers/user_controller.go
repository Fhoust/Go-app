package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	api "github.com/zerepl/go-app/internal/domain/services"
	"github.com/zerepl/go-app/internal/model"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	userService api.UserService
}

func NewUserController(userService api.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (c UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, err.Error())
	}

	user, err := c.userService.GetUser(ctx, int64(id))

	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.String(http.StatusOK, fmt.Sprint(user))
}

func (c UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers(ctx)

	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.String(http.StatusOK, fmt.Sprint(users))
}

func (c UserController) CreateNewUser(ctx *gin.Context) {
	var requestBody model.User

	err := ctx.BindJSON(&requestBody)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, err.Error())
	}

	id, err := c.userService.CreateNewUser(ctx, requestBody)

	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.String(http.StatusOK, strconv.Itoa(int(id)))
}

func (c UserController) UpdateUser(ctx *gin.Context) {
	var requestBody model.User

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, err.Error())
	}

	requestBody.ID = int64(id)

	err = ctx.BindJSON(&requestBody)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, err.Error())
	}

	err = c.userService.UpdateUser(ctx, requestBody)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.String(http.StatusOK, "Ok")
}

func (c UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, err.Error())
	}

	err = c.userService.DeleteUser(ctx, int64(id))
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.String(http.StatusOK, "Ok")
}

func (c UserController) UserRoutes(router *gin.Engine) {

	router.GET("/users", c.GetAllUsers)
	router.GET("/users/:id", c.GetUser)
	router.DELETE("/users/:id", c.DeleteUser)
	router.PUT("/users/:id", c.UpdateUser)
	router.POST("/users", c.CreateNewUser)
}
