package routes

import (
	"github.com/cnAndreLee/tods_server/controller"
	"github.com/gin-gonic/gin"
)

func GetCategoryRoutes(route *gin.RouterGroup) {

	r := route.Group("/categories")

	// category 的增删改查
	r.Handle("POST", "/", controller.CreateCategory)
	r.Handle("DELETE", "/", controller.DeleteCategory)
	r.Handle("PUT", "/", controller.ModifyCategory)
	r.Handle("GET", "/", controller.GetCategory)

}
