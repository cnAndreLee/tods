package controller

import (
	"strings"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/cnAndreLee/tods_server/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func UploadFile(ctx *gin.Context) {

	res := response.ResponseStruct{
		HttpStatus: 200,
		Code:       response.SuccessCode,
		Msg:        "OK",
		Data:       nil,
	}

	// 验证分类名
	fileBelong := ctx.PostForm("filebelong")
	if fileBelong == "" {
		res = response.ResponseStruct{
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
		res = response.ResponseStruct{
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

	suffix := arr[len(arr)-1]
	fileID := uuid.NewV4().String()
	filename := fileID + "." + suffix

	if err := ctx.SaveUploadedFile(file, "./file/"+filename); err != nil {
		utils.LogERROR(err.Error())
		response.ResponseServerError(ctx, "服务器错误，上传失败")
		return
	}

	modelFile := model.File{
		FileID:     fileID,
		FileBelong: fileBelong,
		Title:      arr[0],
		Suffix:     suffix,
	}

	common.DB.Debug().Create(&modelFile)

	res = response.ResponseStruct{
		HttpStatus: 200,
		Code:       response.SuccessCode,
		Msg:        "OK",
		Data:       nil,
	}

	response.Response(ctx, res)
}

func GetFile(ctx *gin.Context) {
	// var modelFiles []model.File

	// common.DB.Find(&modelFiles)

	// var modelResFiles []model.ResFile
	// for _, v := range modelFiles {

	// 	// 拼接文件存储地址
	// 	url := config.CONFIG.SCHEME + "://" + config.CONFIG.ExportHOST + ":" + strconv.Itoa(config.CONFIG.ExportPort) + "/api/files/" + v.FileID + "." + v.Suffix
	// 	fmt.Println(url)
	// 	modelResFiles = append(modelResFiles, model.ResFile{
	// 		FileID:     v.FileID,
	// 		FileBelong: v.FileBelong,
	// 		Title:      v.Title,
	// 		Suffix:     v.Suffix,
	// 		Url:        url,
	// 	})
	// }

	// res := response.ResponseStruct{
	// 	HttpStatus: 200,
	// 	Code:       response.SuccessCode,
	// 	Msg:        "OK",
	// 	Data:       gin.H{"files": modelResFiles},
	// }
	// response.Response(ctx, res)
}
