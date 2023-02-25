package config

import (
	"os"
	"strconv"
)

type ConfigStruct struct {
	IsDebug    bool     `yaml:"IsDebug"`
	PORT       int      `yaml:"PORT"`
	ExportHOST string   `yaml:"ExportHOST"`
	ExportPort int      `yaml:"ExportPort"`
	SCHEME     string   `yaml:"SCHEME"`
	DB         DbStruct `yaml:"DB"`
}

type DbStruct struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
}

var CONFIG ConfigStruct

// 被禁用用户名map
var BannedAccountsMap map[string]string

func InitConfig() {
	// default config
	CONFIG = ConfigStruct{
		IsDebug:    false,
		PORT:       8000,
		ExportHOST: "10.0.0.10",
		ExportPort: 8000,
		DB: DbStruct{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			DbName:   "tods",
		},
	}

	if s, exists := os.LookupEnv("TODS_IsDebug"); exists {
		if s == "true" || s == "True" {
			CONFIG.IsDebug = true
		}
	}

	if s, exists := os.LookupEnv("TODS_PORT"); exists {
		res, err := strconv.Atoi(s)
		if err == nil {
			CONFIG.PORT = res
		}
	}

	if s, exists := os.LookupEnv("TODS_EXPORTHOST"); exists {
		CONFIG.ExportHOST = s
	}
	if s, exists := os.LookupEnv("TODS_EXPORTPORT"); exists {
		res, err := strconv.Atoi(s)
		if err == nil {
			CONFIG.ExportPort = res
		}
	}

	if s, exists := os.LookupEnv("TODS_DB_HOST"); exists {
		CONFIG.DB.Host = s
	}
	if s, exists := os.LookupEnv("TODS_DB_PORT"); exists {
		res, err := strconv.Atoi(s)
		if err == nil {
			CONFIG.DB.Port = res
		}
	}
	if s, exists := os.LookupEnv("TODS_DB_USER"); exists {
		CONFIG.DB.User = s
	}
	if s, exists := os.LookupEnv("TODS_DB_PASSWORD"); exists {
		CONFIG.DB.Password = s
	}
	if s, exists := os.LookupEnv("TODS_DB_DBNAME"); exists {
		CONFIG.DB.DbName = s
	}

	// content, _ := ioutil.ReadFile("config/config.yaml")

	// err := yaml.Unmarshal(content, &CONFIG)

	// if err != nil {
	// 	panic("failed to load config, err: " + err.Error())
	// }

	// if s := os.Getenv("IsDebug"); s == "true" {
	// 	CONFIG.IsDebug = true
	// } else {
	// 	CONFIG.IsDebug = false
	// }

	//载入禁用用户名
	// 非法用户名数组
	BannedAccounts := []string{"root", "admin", "administrator", "account"}
	BannedAccountsMap = make(map[string]string)
	for _, v := range BannedAccounts {
		BannedAccountsMap[v] = ""
	}

	// utils.LogINFO(fmt.Sprintf("config loaded --- %v \n", CONFIG))
}
