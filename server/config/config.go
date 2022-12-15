package config

import (
	"fmt"
	"io/ioutil"

	"github.com/cnAndreLee/tods_server/utils"
	"gopkg.in/yaml.v2"
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
	content, _ := ioutil.ReadFile("config/config.yaml")

	CONFIG = ConfigStruct{}

	err := yaml.Unmarshal(content, &CONFIG)

	if err != nil {
		panic("failed to load config, err: " + err.Error())
	}

	//载入禁用用户名
	// 非法用户名数组
	BannedAccounts := []string{"root", "admin", "administrator", "account"}
	BannedAccountsMap = make(map[string]string)
	for _, v := range BannedAccounts {
		BannedAccountsMap[v] = ""
	}

	utils.LogINFO(fmt.Sprintf("config loaded --- %v \n", CONFIG))
}
