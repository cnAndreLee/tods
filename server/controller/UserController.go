package controller

import (
	"fmt"
	"net/http"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/config"
	"github.com/cnAndreLee/tods_server/dto"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/cnAndreLee/tods_server/service"
	"github.com/cnAndreLee/tods_server/utils"
	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(c *gin.Context) {

	var RequestUser = dto.UserRegistry{}
	//请求结果写入
	c.Bind(&RequestUser)
	utils.LogINFO(fmt.Sprintf("收到注册请求，user: %+v", RequestUser))

	//校验用户名
	if !IsAccountLegal(RequestUser.Account) {
		res := response.ResponseStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       1,
			Msg:        "用户名不合法，请使用小写字母和数字组合，且是字母开头",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	// 数据交给service处理
	res := service.UserRegistryService(RequestUser)

	response.Response(c, res)

}

func JWTLogin(c *gin.Context) {

	var RequestUser = dto.UserLogin{}
	c.Bind(&RequestUser)

	utils.LogINFO(fmt.Sprintf("收到用户登录请求，accout:%v   key:%v", RequestUser.Account, RequestUser.Key))

	// 交给service验证用户
	res := service.UserLoginService(RequestUser)
	response.Response(c, res)
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	var muser model.User
	result := common.DB.Where("account = ?", user).First(&muser)

	// 如果数据库中未查询到该用户，则返回未认证
	if result.Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusUnauthorized,
			Code:       1,
			Msg:        "用户不存在",
			Data:       gin.H{"user": user},
		}
		response.Response(ctx, res)
		return
	}

	// 返回前将密码置为空
	muser.Key = ""
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       0,
		Msg:        "ok",
		Data:       gin.H{"user": muser},
	}

	response.Response(ctx, res)
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
