package routes

import (
	"gin-mongo-api/controllers"
	"gin-mongo-api/services"
	"github.com/gin-gonic/gin"
	)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.POST("/", uc.userController.CreateUser)

}