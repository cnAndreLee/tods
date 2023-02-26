package common

import (
	"time"

	"github.com/cnAndreLee/tods_server/config"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"strconv"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := config.CONFIG.DB.Host
	port := config.CONFIG.DB.Port
	user := config.CONFIG.DB.User
	password := config.CONFIG.DB.Password
	dbname := config.CONFIG.DB.DbName

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + strconv.Itoa(port) + " sslmode=disable"
	utils.LogINFO("dsn is --- " + dsn)

	var db *gorm.DB
	var err error
	for {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			utils.LogERROR("Failed to connect database, err: " + err.Error())
			time.Sleep(time.Duration(5) * time.Second)
		} else {
			utils.LogINFO("Success to connect database!")
			break
		}

	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.File{})

	if config.CONFIG.IsDebug {
		DB = db.Debug()
	} else {
		DB = db
	}

	return db

}

func GetDB() *gorm.DB {
	return DB
}
