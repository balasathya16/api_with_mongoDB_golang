package controllers

import (
	"github.com/gin-gonic/gin"
	"userapi.com/services"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "")

}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")

}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")

}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")

}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/create", uc.DeleteUser)

}
