package routes

import (
	"github.com/cnAndreLee/tods_server/controller"
	"github.com/cnAndreLee/tods_server/middleware"
	"github.com/gin-gonic/gin"
)

func GetUserRoutes(route *gin.RouterGroup) {

	user := route.Group("/user")

	user.Handle("POST", "/login", middleware.ClientIPDeal(), controller.JWTLogin)
	user.Handle("POST", "/register", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.Register)
	user.Handle("GET", "/info", middleware.AuthMiddleware(), controller.Info)

	// 返回所有用户信息列表
	user.Handle("GET", "/users", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.RespUsers)
	user.Handle("DELETE", "", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.DeleteUser)
	user.Handle("GET", "/school", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.RespUsersSchool)

}
