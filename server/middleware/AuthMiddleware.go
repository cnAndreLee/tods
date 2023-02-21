package middleware

import (
	"net/http"
	"strings"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/cnAndreLee/tods_server/utils"
	"github.com/gin-gonic/gin"
)

//AuthMiddleware 判断用户是否登录
func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			res := response.ResponseStruct{
				HttpStatus: http.StatusUnauthorized,
				Code:       response.UnauthorizedCode,
				Data:       nil,
				Msg:        "认证失败",
			}
			response.Response(ctx, res)
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		user, err := common.ParseToken(tokenString)
		if err != nil {
			res := response.ResponseStruct{
				HttpStatus: http.StatusUnauthorized,
				Code:       response.UnauthorizedCode,
				Data:       nil,
				Msg:        "认证失败",
			}
			response.Response(ctx, res)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

// 判断认证用户是否为管理员
func AdminAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		user, _ := ctx.Get("user")

		var doUser model.User
		common.DB.Where("account = ?", user.(string)).First(&doUser)

		if doUser.IsAdmin == false {
			res := response.ResponseStruct{
				HttpStatus: http.StatusForbidden,
				Code:       response.ForbiddenCode,
				Data:       nil,
				Msg:        "无权限",
			}
			response.Response(ctx, res)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func ClientIPDeal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.LogINFO("IP: " + ctx.ClientIP())
		ctx.Next()
	}
}
