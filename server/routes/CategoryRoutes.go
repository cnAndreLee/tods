package routes

import (
	"github.com/cnAndreLee/tods_server/controller"
	"github.com/cnAndreLee/tods_server/middleware"
	"github.com/gin-gonic/gin"
)

func GetCategoryRoutes(route *gin.RouterGroup) {

	r := route.Group("/categories")

	// category 的增删改查
	r.Handle("POST", "/", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.CreateCategory)
	r.Handle("DELETE", "/", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.DeleteCategory)
	r.Handle("PUT", "/", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.ModifyCategory)
	r.Handle("GET", "/", middleware.AuthMiddleware(), middleware.AuthVip(), controller.GetCategory)

	// 更改一级category的权限
	r.Handle("PUT", "/permission", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.ChangeCategoryPermission)
}
