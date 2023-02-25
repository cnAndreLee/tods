package service

import (
	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/config"
	"github.com/cnAndreLee/tods_server/model"
)

// 判断用户是否存在
func IsAccountExist(account string, doUser *model.User) bool {

	result := common.DB.Where("account = ?", account).First(&doUser)

	if result.Error != nil {
		return false
	}

	return true
}

// 判断用户名是否合法
func IsAccountLegal(account string) bool {
	// LigalChars := "abcdefghijkmlnopqrstuvwxyz0123456789"

	// 判断字符串是否在a-z, 0-9范围之内
	for k, v := range account {

		if k == 0 {
			if !(v >= 97 && v <= 122) {
				return false
			}
		}

		if !((v >= 48 && v <= 57) || (v >= 97 && v <= 122)) {
			return false
		}
	}

	if _, ok := config.BannedAccountsMap[account]; ok {
		return false
	}

	return true

}
