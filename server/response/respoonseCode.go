package response

import "github.com/gin-gonic/gin"

type ResponseStruct struct {
	HttpStatus int    //http状态
	Code       int    //状态码
	Data       gin.H  //数据
	Msg        string //信息
}

const (
	SuccessCode      = 2000
	FailCode         = 4000
	CheckFailCode    = 4022
	ServerErrorCode  = 5000
	TokenExpriedCode = 4010 //token过期
	UnauthorizedCode = 4011 //未授权
	ForbiddenCode    = 4030 //无管理员权限
	VipExiredCode    = 4031 //会员过期

	CategoryNameExistedCode          = 4100
	SecondaryCategoryNameExistedCode = 4100
)
