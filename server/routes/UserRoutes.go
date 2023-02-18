package routes

import (
	"github.com/cnAndreLee/tods_server/controller"
	"github.com/cnAndreLee/tods_server/middleware"
	"github.com/gin-gonic/gin"
)

func GetUserRoutes(route *gin.RouterGroup) {

	user := route.Group("/user")

	user.Handle("POST", "/login", middleware.ClientIPDeal(), controller.JWTLogin)
	user.Handle("POST", "/register", controller.Register)
	user.Handle("GET", "/info", middleware.AuthMiddleware(), controller.Info)

	user.Handle("GET", "/users", middleware.AuthMiddleware(), controller.RespUsers)

}
