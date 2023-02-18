package service

import (
	"net/http"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/config"
	"github.com/cnAndreLee/tods_server/dto"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/cnAndreLee/tods_server/utils"
)

func UserRegistryService(user dto.UserRegistry) response.ResponseStruct {

	var res response.ResponseStruct

	// muser: 接收数据库结果
	var muser model.User

	// 判断用户名是否已存在
	exists := IsAccountExist(user.Account, &muser)

	if exists {
		utils.LogINFO("注册失败，用户已存在")
		res = response.ResponseStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Data:       nil,
			Msg:        "注册失败，用户已存在",
		}

		return res
	}

	NewUser := model.User{
		Account:    user.Account,
		Key:        user.Key,
		Class:      user.Class,
		Schoolname: user.Schoolname,
		Telephone:  user.Telephone,
		Outtime:    user.Outtime,
	}

	//用户数据写入
	common.DB.Create(&NewUser)
	res = response.ResponseStruct{
		HttpStatus: http.StatusBadRequest,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        "注册成功",
	}

	return res

}

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
				utils.LogINFO("字符串不合法: 非小写字母开头")
				return false
			}
		}

		if !((v >= 48 && v <= 57) || (v >= 97 && v <= 122)) {
			utils.LogINFO("字符串不合法")
			return false
		}
	}

	if _, ok := config.BannedAccountsMap[account]; ok {
		utils.LogINFO("用户名被禁止")
		return false
	}

	return true

}
