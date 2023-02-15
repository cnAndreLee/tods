package common

// import (
//     "database/sql"
//     "fmt"

//     _ "github.com/lib/pq"
// )

// var DB *sql.DB

// func InitDB() {
//     psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmoe=disable", "localhost", 54321, "postgres", "postgres", "users")
//     db, err := sql.Open("postgres", psqlInfo)

//     if err != nil {
//         return
//     }

//     DB = db
// }

import (
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
	//dsn := "host=localhost user=postgres password=postgres dbname=tods port=54321 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
	// db.AutoMigrate(&model.SecondaryCategory{})
	db.AutoMigrate(&model.File{})

	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB
}
