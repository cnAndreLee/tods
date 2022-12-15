package controller

import (
	"fmt"
	"net/http"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/cnAndreLee/tods_server/utils"
	"github.com/cnAndreLee/tods_server/vo"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func CreateCategory(c *gin.Context) {

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data:       nil,
	}

	var voCategory vo.CreateCategory
	if err := c.ShouldBind(&voCategory); err != nil {
		res = response.ResponseStruct{
			HttpStatus: http.StatusBadRequest,
			Code:       response.FailCode,
			Msg:        "数据无效",
			Data:       nil,
		}
		utils.LogINFO(fmt.Sprintf("创建分类失败，Category: ", voCategory))
		response.Response(c, res)
		return
	}
	// 生成UUID
	voCategory.ID = uuid.NewV4().String()
	utils.LogINFO(fmt.Sprintf("创建分类，Category: ", voCategory))

	// 写入数据库
	modelCategory := model.Category{ID: voCategory.ID, Name: voCategory.Name}
	result := common.DB.Debug().Create(modelCategory)
	if err := result.Error; err != nil {

		res = response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.CategoryNameExistedCode,
			Msg:        "写入失败，重名",
			Data:       nil,
		}
	}

	response.Response(c, res)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	utils.LogINFO(fmt.Sprint("收到删除分类请求，分类id：", id))

	modelCategory := model.Category{ID: id}
	result := common.DB.Debug().Delete(&modelCategory)
	// 不会进入此逻辑，因为删除一直成功
	if err := result.Error; err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "删除失败",
			Data:       nil,
		}
		response.Response(c, res)
		utils.LogINFO(err.Error())
		return
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "ok",
		Data:       nil,
	}

	response.Response(c, res)
}

func ModifyCategory(c *gin.Context) {
	id := c.Param("id")
	modelCategory := model.Category{ID: id}
	c.Bind(&modelCategory)

	utils.LogINFO(fmt.Sprint("修改分类，modelCategory:", modelCategory))

	if modelCategory.Name == "" {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "修改失败，名称不合法",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	result := common.DB.Debug().Model(&modelCategory).Where("id = ?", id).Update("name", modelCategory.Name)
	if err := result.Error; err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "修改失败",
			Data:       nil,
		}
		response.Response(c, res)
		utils.LogINFO(err.Error())
		return
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "ok",
		Data:       nil,
	}

	response.Response(c, res)
}

func GetCategory(c *gin.Context) {
	var modelCategories []model.Category
	common.DB.Find(&modelCategories)

	utils.LogINFO(fmt.Sprint(modelCategories))

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "ok",
		Data: gin.H{
			"Categories": modelCategories,
		},
	}

	response.Response(c, res)
}
