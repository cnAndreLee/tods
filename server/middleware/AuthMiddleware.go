package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/cnAndreLee/tods_server/utils"
	"github.com/gin-gonic/gin"
)

//AuthMiddleware 判断用户是否登录
func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		fmt.Printf("token is %v\n", tokenString)

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			res := response.ResponseStruct{
				HttpStatus: http.StatusUnauthorized,
				Code:       1,
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
				Code:       1,
				Data:       nil,
				Msg:        "认证失败",
			}
			response.Response(ctx, res)

			ctx.Abort()
			return
		}

		fmt.Printf("token parse success ,user is %v", user)

		ctx.Set("user", user)
		ctx.Next()
	}
}

func ClientIPDeal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.LogINFO("IP: " + ctx.ClientIP())
		ctx.Next()
	}
}
