package main

import (
	"fmt"
	"strconv"

	// "github.com/gin-contrib/cors"
	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/config"
	"github.com/cnAndreLee/tods_server/routes"
	"github.com/cnAndreLee/tods_server/utils"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/postgres"
)

type UserInfo struct {
	accout   string
	password string
}

type AuthResult struct {
	token string
	msg   string
}

func main() {
	fmt.Println("main start")
	utils.LogINFO("main start")
	config.InitConfig()

	db := common.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := gin.Default()

	//解决跨域问题
	// router.Use(cors.New(cors.Config{
	//     AllowOrigins:     []string{"http://10.0.0.137:8080"},
	//     AllowMethods:     []string{"GET", "POST"},
	//     AllowHeaders:     []string{"Origin"},
	//     ExposeHeaders:    []string{"Content-Length"},
	//     AllowCredentials: true,
	//     AllowOriginFunc: func(origin string) bool {
	//         return origin == "https://github.com"
	//     },
	//     MaxAge: 12 * time.Hour,
	// }))

	r = routes.CollectRoute(r)

	port := config.CONFIG.PORT
	panic(r.Run(":" + strconv.Itoa(port)))

}
