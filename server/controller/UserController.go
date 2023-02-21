package controller

import (
	"fmt"
	"net/http"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/cnAndreLee/tods_server/service"
	"github.com/cnAndreLee/tods_server/utils"
	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(c *gin.Context) {

	var dtoUser model.DtoUserRegistry
	if err := c.ShouldBind(&dtoUser); err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "注册失败，参数错误",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	//校验用户名
	if !service.IsAccountLegal(dtoUser.Account) {
		res := response.ResponseStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Msg:        "用户名不合法，请使用小写字母和数字组合，且是字母开头",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	NewUser := model.User{
		Account: dtoUser.Account,
		Key:     dtoUser.Key,
		IsAdmin: dtoUser.IsAdmin,
		School:  dtoUser.School,
		OutDate: dtoUser.OutDate,
		Remark:  dtoUser.Remark,
	}
	//用户数据写入
	if common.DB.Create(&NewUser).Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Data:       nil,
			Msg:        "注册失败，用户已存在",
		}
		response.Response(c, res)
		return
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        "注册成功",
	}
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

	var dtoUserLogin model.DtoUserLogin
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

func DeleteUser(ctx *gin.Context) {

	account := ctx.Query("account")

	if account == "" {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "删除失败",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	result := common.DB.Where("account = ?", account).Delete(&model.User{})
	if result.Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.ServerErrorCode,
			Msg:        "删除失败, 服务器错误",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "OK",
		Data:       nil,
	}
	response.Response(ctx, res)

}

// 响应所有学校
func RespUsersSchool(ctx *gin.Context) {

	var schools []string
	common.DB.Model(&model.User{}).Distinct().Pluck("school", &schools)

	var filteredSchools []string
	for _, v := range schools {
		if v != "" {
			filteredSchools = append(filteredSchools, v)
		}
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data: gin.H{
			"schools": filteredSchools,
		},
	}
	response.Response(ctx, res)
}
