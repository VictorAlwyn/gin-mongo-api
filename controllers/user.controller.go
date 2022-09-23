package controllers

import (
	"context"
    "net/http"

	"github.com/gin-gonic/gin"
	"gin-mongo-api/models"
	"gin-mongo-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	userService services.UserService
	ctx         context.Context
	collection  *mongo.Collection
}

func NewUserController(userService services.UserService, ctx context.Context, collection *mongo.Collection) UserController {
	return UserController{userService, ctx, collection}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user *models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := uc.userService.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUser})
}