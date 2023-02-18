package main

import (
	"strconv"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/config"
	"github.com/cnAndreLee/tods_server/routes"
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

	config.InitConfig()

	db := common.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if !config.CONFIG.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r = routes.CollectRoute(r)

	port := config.CONFIG.PORT
	panic(r.Run(":" + strconv.Itoa(port)))

}
