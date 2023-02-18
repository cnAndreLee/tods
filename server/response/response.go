package response

import (
	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, res ResponseStruct) {
	ctx.JSON(res.HttpStatus, gin.H{"status": res.Code, "msg": res.Msg, "data": res.Data})
}

func ResponseRequestError(ctx *gin.Context, str string) {
	ctx.JSON(400, gin.H{"status": FailCode, "msg": str, "data": nil})
}

func ResponseServerError(ctx *gin.Context, str string) {
	ctx.JSON(500, gin.H{"status": ServerErrorCode, "msg": str, "data": nil})

}

// func Success(ctx *gin.Context, code int, data gin.H, msg string) {
//     Response(ctx, http.StatusOK, 200, data, msg)
// }

// func Fail(ctx *gin.Context, code int, data gin.H, msg string) {
//     Response(ctx, http.StatusOK, 400, data, msg)
// }
