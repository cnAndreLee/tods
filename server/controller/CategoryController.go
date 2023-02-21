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
	uuid "github.com/satori/go.uuid"
)

func CreateCategory(c *gin.Context) {

	// 接收客户端DTO
	var dtoCreateCategory model.CreateCategoryDTO
	if err := c.ShouldBind(&dtoCreateCategory); err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "数据无效",
			Data:       nil,
		}

		response.Response(c, res)
		return
	}

	// 转换为DO
	var doCategory model.Category
	doCategory.ID = uuid.NewV4().String() // 生成UUID
	doCategory.ParentID = dtoCreateCategory.ParentID
	if dtoCreateCategory.Level == 1 {
		doCategory.ParentID = "0"
	}
	doCategory.Level = dtoCreateCategory.Level
	doCategory.Title = dtoCreateCategory.Title
	doCategory.Permissions = nil

	// 写入数据库之前判断parentid是否存在
	if doCategory.Level > 1 {
		// 如果查找错误，证明找不到对应parentid
		if common.DB.Where("id = ?", doCategory.ParentID).First(&model.Category{}).Error != nil {
			res := response.ResponseStruct{
				HttpStatus: http.StatusBadRequest,
				Code:       response.FailCode,
				Msg:        "数据无效,不存在对应parentid",
				Data:       nil,
			}
			utils.LogINFO("不存在对应parentid")
			response.Response(c, res)
			return
		}
	}

	// 写入数据库
	result := common.DB.Create(doCategory)
	if err := result.Error; err != nil {

		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.CategoryNameExistedCode,
			Msg:        "写入失败，重名",
			Data:       nil,
		}
		utils.LogINFO("写入失败，重名")
		response.Response(c, res)
		return
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data:       nil,
	}

	response.Response(c, res)
	utils.LogINFO(fmt.Sprintf("----------分类创建成功--------"))
}

func DeleteCategory(c *gin.Context) {

	// 接收客户端DTO
	var dtoDeleteCategory model.DeleteCategoryDTO
	if err := c.ShouldBind(&dtoDeleteCategory); err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "数据无效",
			Data:       nil,
		}
		utils.LogINFO(fmt.Sprintf("删除分类失败, 数据无效, dtoDeleteCategory: ", dtoDeleteCategory))
		response.Response(c, res)
		return
	}

	var doCategory model.Category
	doCategory.ID = dtoDeleteCategory.ID

	// 判断是否存在id，不存在则返回错误
	if common.DB.First(&doCategory).Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "不存在此ID",
			Data:       nil,
		}
		utils.LogINFO("no record")
		response.Response(c, res)
		return
	}

	// 查找child，存在child则返回错误 ()
	if service.IsExistChild(&doCategory) {
		msg := ""
		if doCategory.Level == 1 {
			msg = "删除失败, 分类下存在合集"
		} else {
			msg = "删除失败, 合集下存在资源"
		}
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        msg,
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	// 删除操作
	if common.DB.Delete(&doCategory).Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "删除失败, database error",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "OK",
		Data:       nil,
	}
	response.Response(c, res)

	user, _ := c.Get("user")
	logMsg := "----------分类删除成功, 操作用户:" + fmt.Sprint(user) + " 删除的分类:" + doCategory.Title + "--------"
	utils.LogINFO(logMsg)
}

// change title
func ModifyCategory(c *gin.Context) {

	// 接收DTO
	var dtoModifyCategory model.ModifyCategoryDTO
	if err := c.ShouldBind(&dtoModifyCategory); err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "参数错误",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	// 检查分类id是否存在
	if common.DB.Where("id = ?", dtoModifyCategory.ID).First(&model.Category{}).Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "修改失败, 不存在此id",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	result := common.DB.
		Model(&model.Category{}).
		Where("id = ?", dtoModifyCategory.ID).
		Update("title", dtoModifyCategory.NewTitle)

	if err := result.Error; err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "修改失败, title已存在",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data:       nil,
	}
	response.Response(c, res)

}

// change permissions
func ChangeCategoryPermission(c *gin.Context) {

	// 接收DTO
	var dtoChangePermission model.ChangeCategoryPermissionDTO
	if err := c.ShouldBind(&dtoChangePermission); err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "参数错误",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	// 检查分类id是否存在
	if common.DB.Where("id = ?", dtoChangePermission.ID).First(&model.Category{}).Error != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "修改失败, 不存在此id",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	fmt.Print(dtoChangePermission)
	result := common.DB.
		Model(&model.Category{}).
		Where("id = ?", dtoChangePermission.ID).
		Update("permissions", dtoChangePermission.NewPermissions)

	if err := result.Error; err != nil {
		res := response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "修改失败, 修改时发生错误",
			Data:       nil,
		}
		response.Response(c, res)
		return
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data:       nil,
	}
	response.Response(c, res)

}

func GetCategory(c *gin.Context) {

	var Categories []model.Category
	common.DB.Find(&Categories)

	var voGetCategories []model.GetCategoryVO
	for _, v := range Categories {
		if v.Level == 1 {
			voGetCategories = append(voGetCategories, model.GetCategoryVO{
				ID:          v.ID,
				Title:       v.Title,
				Children:    nil, //[]model.GetCategoryVO{},
				Permissions: v.Permissions,
			})
			for _, v2 := range Categories {
				if v2.Level == 2 {
					if v2.ParentID == v.ID {
						voGetCategories[len(voGetCategories)-1].Children = append(voGetCategories[len(voGetCategories)-1].Children, model.GetCategoryVO{
							ID:       v2.ID,
							Title:    v2.Title,
							Children: nil,
						})
					}
				}
			}
		}
	}

	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data: gin.H{
			"categories": voGetCategories,
		},
	}
	response.Response(c, res)
}
