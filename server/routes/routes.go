package routes

import (
	"net/http"
	"os"

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
		GetFIleRoutes(v1)
	}

	if _, err := os.Stat("./file"); err != nil {
		panic("ERROR!!!!! file not exist")
	}
	r.StaticFS("/api/files", http.Dir("./file/"))

	return r

}
