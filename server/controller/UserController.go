package controller

import (
	"fmt"
	"net/http"

	"github.com/cnAndreLee/tods_server/common"
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
	if !service.IsAccountLegal(RequestUser.Account) {
		res := response.ResponseStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
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

func JWTLogin(ctx *gin.Context) {

	// 定义响应结构体
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data:       nil,
	}

	var dtoUserLogin dto.UserLogin
	if err := ctx.ShouldBind(&dtoUserLogin); err != nil {
		res = response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "登录失败，参数错误",
			Data:       nil,
		}
		response.Response(ctx, res)
		utils.LogINFO("登录失败，参数错误:" + err.Error())
		return
	}

	utils.LogINFO(fmt.Sprintf("收到用户登录请求，accout:%v   key:%v", dtoUserLogin.Account, dtoUserLogin.Key))

	// 交给service验证用户
	// 验证用户是否存在，并关联do与数据库
	var doUser model.User
	exists := service.IsAccountExist(dtoUserLogin.Account, &doUser)

	if !exists {
		utils.LogINFO("登录失败，用户不存在")
		res = response.ResponseStruct{
			HttpStatus: 200,
			Code:       response.FailCode,
			Data:       nil,
			Msg:        "登录失败，用户名或密码错误",
		}
		response.Response(ctx, res)
		return
	}

	// 验证密码
	if dtoUserLogin.Key != doUser.Key {
		utils.LogINFO("登录失败，密码不匹配")
		res = response.ResponseStruct{
			HttpStatus: 200,
			Code:       1,
			Data:       nil,
			Msg:        "登录失败，用户名或密码错误",
		}
		response.Response(ctx, res)
		return
	}

	// 生成token
	tokenString, err := common.CreateToken(doUser.Account)
	if err != nil {
		res = response.ResponseStruct{
			HttpStatus: http.StatusInternalServerError,
			Code:       response.ServerErrorCode,
			Data:       nil,
			Msg:        "登录失败，服务器错误",
		}
		response.Response(ctx, res)
		return
	}

	//return token
	res = response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"token": tokenString},
		Msg:        "ok",
	}

	response.Response(ctx, res)
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	var muser model.User
	result := common.DB.Where("account = ?", user).First(&muser)

	// 如果数据库中未查询到该用户，则返回未认证
	if result.Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusUnauthorized,
			Code:       response.FailCode,
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
		Code:       response.SuccessCode,
		Msg:        "ok",
		Data:       gin.H{"user": muser},
	}

	response.Response(ctx, res)
}

// 响应所有用户信息表
func RespUsers(ctx *gin.Context) {

	// 获取用户名，判断是否为admin
	user, _ := ctx.Get("user")
	var doUser model.User
	result := common.DB.Where("account = ?", user).First(&doUser)
	// 如果数据库中未查询到该用户，则返回未认证
	if result.Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusUnauthorized,
			Code:       1,
			Msg:        "用户不存在",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	if doUser.Class != "admin" {
		res := response.ResponseStruct{
			HttpStatus: http.StatusForbidden,
			Code:       1,
			Msg:        "用户无权限",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	var users []model.User
	common.DB.Find(&users)

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data: gin.H{
			"users": users,
		},
	}
	response.Response(ctx, res)
}
