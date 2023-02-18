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

	utils.LogINFO(fmt.Sprintf("----------收到分类创建请求--------"))

	// 定义响应结构体
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data:       nil,
	}

	// 接收客户端DTO
	var dtoCreateCategory model.CreateCategoryDTO
	if err := c.ShouldBind(&dtoCreateCategory); err != nil {
		res = response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "数据无效",
			Data:       nil,
		}
		utils.LogINFO(fmt.Sprintf("创建分类失败, 数据无效, voCreateCategory: ", dtoCreateCategory))
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

	// 写入数据库之前判断parentid是否存在
	if doCategory.Level > 1 {
		// 如果查找错误，证明找不到对应parentid
		if common.DB.Where("id = ?", doCategory.ParentID).First(&model.Category{}).Error != nil {
			res = response.ResponseStruct{
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

		res = response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.CategoryNameExistedCode,
			Msg:        "写入失败，重名",
			Data:       nil,
		}
		utils.LogINFO("写入失败，重名")
		response.Response(c, res)
		return
	}

	response.Response(c, res)
	utils.LogINFO(fmt.Sprintf("----------分类创建成功--------"))
}

func DeleteCategory(c *gin.Context) {

	utils.LogINFO(fmt.Sprintf("----------收到分类删除请求--------"))

	// 定义响应结构体
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data:       nil,
	}

	// 接收客户端DTO
	var dtoDeleteCategory model.DeleteCategoryDTO
	if err := c.ShouldBind(&dtoDeleteCategory); err != nil {
		res = response.ResponseStruct{
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
		res = response.ResponseStruct{
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
		res = response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "删除失败,存在child",
			Data:       nil,
		}
		response.Response(c, res)
		utils.LogINFO("删除失败,存在child")
		return
	}

	result := common.DB.Delete(&doCategory)
	// 不会进入此逻辑，因为删除一直成功
	if result.Error != nil {
		res = response.ResponseStruct{
			HttpStatus: http.StatusOK,
			Code:       response.FailCode,
			Msg:        "删除失败",
			Data:       nil,
		}
		response.Response(c, res)
		utils.LogINFO("删除失败, Error: " + result.Error.Error())
		return
	}

	response.Response(c, res)
	utils.LogINFO(fmt.Sprintf("----------分类删除成功--------"))
}

func ModifyCategory(c *gin.Context) {

	utils.LogINFO(fmt.Sprintf("^^^^^^^^^^收到分类修改请求^^^^^^^^^^"))

	// 定义响应结构体
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data:       nil,
	}

	// 接收DTO
	var dtoModifyCategory model.ModifyCategoryDTO
	if err := c.ShouldBind(&dtoModifyCategory); err != nil {
		res = response.ResponseStruct{
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
		utils.LogINFO("修改失败, title已存在,ERROR: " + err.Error())
		return
	}

	response.Response(c, res)
	utils.LogINFO(fmt.Sprintf("vvvvvvvvvvvv分类改名成功vvvvvvvvvvvv"))
}

func GetCategory(c *gin.Context) {

	utils.LogINFO(fmt.Sprintf("^^^^^^^^^^收到分类查询请求^^^^^^^^^^"))

	// 定义响应结构体
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data:       nil,
	}

	// common.DB.Table("categories").
	// 	Select("categories.id, categories.title, categories.id,categories.title").
	// 	Joins("left join categories on categories.parent_id = categories.id")

	var Categories []model.Category
	common.DB.Find(&Categories)
	utils.LogINFO(fmt.Sprint(Categories))

	var voGetCategories []model.GetCategoryVO
	for _, v := range Categories {
		if v.Level == 1 {
			voGetCategories = append(voGetCategories, model.GetCategoryVO{
				ID:       v.ID,
				Title:    v.Title,
				Children: nil, //[]model.GetCategoryVO{},
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

	res = response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Msg:        "",
		Data: gin.H{
			"categories": voGetCategories,
		},
	}

	utils.LogINFO(fmt.Sprint(voGetCategories))

	response.Response(c, res)

}
