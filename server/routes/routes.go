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

	return r

}
