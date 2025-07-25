package router

import (
	"demo/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controller.UserController) *gin.Engine {
	r := gin.Default()
	userController.RegisterRoutes(r)
	return r
}
