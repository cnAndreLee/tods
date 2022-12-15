package routes

import (
	"net/http"

	"github.com/cnAndreLee/tods_server/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	// 解决跨域问题
	r.Use(middleware.CORSMiddleware())

	v1 := r.Group("/api/v1")

	{
		GetUserRoutes(v1)
		GetCategoryRoutes(v1)
		GetSecondaryCategoryRoutes(v1)
		GetFIleRoutes(v1)
	}

	r.StaticFS("/api/files", http.Dir("./file/"))

	// r.Handle("POST", "/api/v1/token", controller.JWTLogin)
	// r.Handle("POST", "/api/v2/register", controller.Register)
	// r.Handle("GET", "/api/v3/info", middleware.AuthMiddleware(), controller.Info)

	//动态路由
	/*
	   r.GET("/:id", func(c *gin.Context) {
	       id := c.Param("id")
	       //可以格式化字符串
	       c.String(200, "get id= %v", id)
	   })
	*/

	return r

}
