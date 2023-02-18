package routes

import (
	"github.com/cnAndreLee/tods_server/controller"
	"github.com/gin-gonic/gin"
)

func GetFIleRoutes(route *gin.RouterGroup) {

	r := route.Group("/file")
	r.Handle("POST", "", controller.UploadFile)
	r.Handle("DELETE", "", controller.DeleteFile)
	// 获取分类ID为id下的文件列表
	r.Handle("GET", "", controller.GetFile)
}
