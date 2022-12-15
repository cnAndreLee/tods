package service

import (
	"net/http"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/dto"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/cnAndreLee/tods_server/utils"
	"github.com/gin-gonic/gin"
)

func UserRegistryService(user dto.UserRegistry) response.ResponseStruct {

	var res response.ResponseStruct

	// muser: 接收数据库结果
	var muser model.User

	// 判断用户名是否已存在
	exists := isAccountExist(user.Account, &muser)

	if exists {
		utils.LogINFO("注册失败，用户已存在")
		res = response.ResponseStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       1,
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
	common.DB.Debug().Create(&NewUser)
	res = response.ResponseStruct{
		HttpStatus: http.StatusBadRequest,
		Code:       0,
		Data:       nil,
		Msg:        "注册成功",
	}

	return res

}

func UserLoginService(user dto.UserLogin) response.ResponseStruct {

	var res response.ResponseStruct

	// muser: 接收数据库结果
	var muser model.User

	exists := isAccountExist(user.Account, &muser)

	if !exists {
		utils.LogINFO("登录失败，用户不存在")
		res = response.ResponseStruct{
			HttpStatus: 200,
			Code:       1,
			Data:       nil,
			Msg:        "登录失败，用户名或密码错误",
		}

		return res
	}

	if muser.Key != user.Key {
		utils.LogINFO("登录失败，密码不匹配")
		res = response.ResponseStruct{
			HttpStatus: 200,
			Code:       1,
			Data:       nil,
			Msg:        "登录失败，用户名或密码错误",
		}

		return res
	}

	// 生成token
	tokenString, err := common.CreateToken(user.Account)
	if err != nil {
		res = response.ResponseStruct{
			HttpStatus: http.StatusInternalServerError,
			Code:       1,
			Data:       nil,
			Msg:        "登录失败，服务器错误",
		}
		return res
	}

	//return token
	res = response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       0,
		Data:       gin.H{"token": tokenString},
		Msg:        "ok",
	}

	return res
}

// 判断用户是否存在
func isAccountExist(account string, muser *model.User) bool {
	result := common.DB.Where("account = ?", account).First(&muser)
	// result := common.DB.First(&user, "account = ?", account)

	if result.Error != nil {
		return false
	}

	return true
}
