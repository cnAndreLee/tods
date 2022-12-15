package routes

import (
	"github.com/cnAndreLee/tods_server/controller"
	"github.com/gin-gonic/gin"
)

func GetFIleRoutes(route *gin.RouterGroup) {

	route.Handle("POST", "/file", controller.UpdaloadFile)
	route.Handle("GET", "/file", controller.GetFile)
}
