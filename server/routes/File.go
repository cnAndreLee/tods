package routes

import (
	"github.com/cnAndreLee/tods_server/controller"
	"github.com/cnAndreLee/tods_server/middleware"
	"github.com/gin-gonic/gin"
)

func GetFIleRoutes(route *gin.RouterGroup) {

	r := route.Group("/file")
	r.Handle("POST", "", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.UploadFile)
	r.Handle("DELETE", "", middleware.AuthMiddleware(), middleware.AdminAuth(), controller.DeleteFile)

	// 获取分类ID为id下的文件列表
	r.Handle("GET", "", middleware.AuthMiddleware(), controller.GetFile)
}
