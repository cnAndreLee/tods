package controller

import (
	"net/http"
	"os"
	"strings"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/cnAndreLee/tods_server/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func UploadFile(ctx *gin.Context) {

	// 验证分类名
	fileBelong := ctx.Query("filebelong")
	if fileBelong == "" {
		res := response.ResponseStruct{
			HttpStatus: 200,
			Code:       response.FailCode,
			Msg:        "分类名错误",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	// 查询是否有对应的分类id
	result := common.DB.Where("id = ?", fileBelong).First(&model.Category{})
	if result.Error != nil {
		res := response.ResponseStruct{
			HttpStatus: 200,
			Code:       response.FailCode,
			Msg:        "分类不存在",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	file, _ := ctx.FormFile("file")
	utils.LogINFO("接收文件：" + file.Filename)

	// 文件名分割
	arr := strings.Split(file.Filename, ".")
	fileTitle := arr[0]
	suffix := arr[len(arr)-1]
	fileID := uuid.NewV4().String()
	filename := fileID + "." + suffix

	// 检查是否存在相同文件名
	if common.DB.Where("title = ?", fileTitle).First(&model.File{}).Error == nil {
		res := response.ResponseStruct{
			HttpStatus: 200,
			Code:       response.FailCode,
			Msg:        "文件名重复",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	if err := ctx.SaveUploadedFile(file, "./file/"+filename); err != nil {
		utils.LogERROR(err.Error())
		response.ResponseServerError(ctx, "服务器错误，上传失败")
		return
	}

	modelFile := model.File{
		FileID:     fileID,
		FileBelong: fileBelong,
		Title:      fileTitle,
		Suffix:     suffix,
	}

	common.DB.Create(&modelFile)

	res := response.ResponseStruct{
		HttpStatus: 200,
		Code:       response.SuccessCode,
		Msg:        "OK",
		Data:       nil,
	}

	response.Response(ctx, res)
}

func DeleteFile(ctx *gin.Context) {

	fileID := ctx.Query("id")

	if fileID == "" {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "删除失败, 参数错误",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	var doFile = model.File{}
	common.DB.Where("file_id = ?", fileID).First(&doFile)

	if common.DB.Where("file_id = ?", fileID).Delete(&model.File{}).Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "删除失败, 数据库错误",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	filePath := "./file/" + doFile.FileID + "." + doFile.Suffix
	if err := os.Remove(filePath); err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "删除失败, 请向开发者报告",
			Data:       nil,
		}
		common.DB.Create(&doFile)
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
	user, _ := ctx.Get("user")
	utils.LogINFO("Deleted File:" + doFile.Title + "  user:" + user.(string))
}

func GetFile(ctx *gin.Context) {

	// 获取地址参数
	parentID := ctx.Query("id")
	if parentID == "" {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "请指定id以查询文件列表",
			Data:       nil,
		}
		response.Response(ctx, res)
		return
	}

	var doFiles []model.File
	common.DB.Where("file_belong = ?", parentID).Find(&doFiles)

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data: gin.H{
			"files": doFiles,
		},
	}
	response.Response(ctx, res)
}
