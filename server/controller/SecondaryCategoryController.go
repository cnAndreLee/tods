package controller

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/cnAndreLee/tods_server/common"
// 	"github.com/cnAndreLee/tods_server/model"
// 	"github.com/cnAndreLee/tods_server/response"
// 	"github.com/cnAndreLee/tods_server/utils"
// 	"github.com/cnAndreLee/tods_server/vo"
// 	"github.com/gin-gonic/gin"
// 	uuid "github.com/satori/go.uuid"
// )

// func CreateSecondaryCategory(c *gin.Context) {

// 	var voSecondaryCategory vo.CreateSecondaryCategory
// 	if err := c.ShouldBind(&voSecondaryCategory); err != nil {
// 		res := response.ResponseStruct{
// 			HttpStatus: http.StatusOK,
// 			Code:       response.FailCode,
// 			Msg:        "数据无效",
// 			Data:       nil,
// 		}
// 		utils.LogINFO(fmt.Sprintf("创建二级分类失败，SecondaryCategory: ", voSecondaryCategory, err))
// 		response.Response(c, res)
// 		return
// 	}

// 	// 查询是否存在对应fatherid
// 	modelCategory := model.Category{ID: voSecondaryCategory.FatherID}
// 	if common.DB.Debug().First(&modelCategory).Error != nil {
// 		res := response.ResponseStruct{
// 			HttpStatus: http.StatusBadRequest,
// 			Code:       response.FailCode,
// 			Msg:        "数据无效,不存在对应fatherid",
// 			Data:       nil,
// 		}
// 		utils.LogINFO("不存在对应fatherid")
// 		response.Response(c, res)
// 		return
// 	}

// 	// 生成UUID
// 	voSecondaryCategory.ID = uuid.NewV4().String()
// 	utils.LogINFO(fmt.Sprintf("创建二级分类，SecondaryCategory: ", voSecondaryCategory))

// 	// 写入数据库
// 	modelSecondaryCategory := model.SecondaryCategory{
// 		ID:       voSecondaryCategory.ID,
// 		Name:     voSecondaryCategory.Name,
// 		FatherID: voSecondaryCategory.FatherID,
// 	}
// 	result := common.DB.Debug().Create(modelSecondaryCategory)
// 	if err := result.Error; err != nil {
// 		res := response.ResponseStruct{
// 			HttpStatus: http.StatusOK,
// 			Code:       response.FailCode,
// 			Msg:        "写入失败",
// 			Data:       nil,
// 		}
// 		response.Response(c, res)
// 		return
// 	}

// 	res := response.ResponseStruct{
// 		HttpStatus: http.StatusOK,
// 		Code:       response.SuccessCode,
// 		Msg:        "",
// 		Data:       nil,
// 	}

// 	response.Response(c, res)
// }

// func DeleteSecondaryCategory(c *gin.Context) {
// 	id := c.Param("id")
// 	utils.LogINFO(fmt.Sprint("收到删除二级分类请求，分类id：", id))

// 	modelSecondaryCategory := model.SecondaryCategory{ID: id}
// 	result := common.DB.Debug().Delete(&modelSecondaryCategory)
// 	// 不会进入此逻辑，因为删除一直成功
// 	if err := result.Error; err != nil {
// 		res := response.ResponseStruct{
// 			HttpStatus: http.StatusOK,
// 			Code:       response.FailCode,
// 			Msg:        "删除失败",
// 			Data:       nil,
// 		}
// 		response.Response(c, res)
// 		utils.LogINFO(err.Error())
// 		return
// 	}

// 	res := response.ResponseStruct{
// 		HttpStatus: http.StatusOK,
// 		Code:       response.SuccessCode,
// 		Msg:        "ok",
// 		Data:       nil,
// 	}

// 	response.Response(c, res)
// }

// func ModifySecondaryCategory(c *gin.Context) {
// 	id := c.Param("id")
// 	modelSecondaryCategory := model.SecondaryCategory{ID: id}
// 	c.Bind(&modelSecondaryCategory)

// 	utils.LogINFO(fmt.Sprint("修改二级分类，modelSecondaryCategory:", modelSecondaryCategory))

// 	if modelSecondaryCategory.Name == "" {
// 		res := response.ResponseStruct{
// 			HttpStatus: http.StatusOK,
// 			Code:       response.FailCode,
// 			Msg:        "修改失败，名称不合法",
// 			Data:       nil,
// 		}
// 		response.Response(c, res)
// 		return
// 	}

// 	result := common.DB.Debug().Model(&modelSecondaryCategory).Where("id = ?", id).Update("name", modelSecondaryCategory.Name)
// 	if err := result.Error; err != nil {
// 		res := response.ResponseStruct{
// 			HttpStatus: http.StatusOK,
// 			Code:       response.FailCode,
// 			Msg:        "修改失败",
// 			Data:       nil,
// 		}
// 		response.Response(c, res)
// 		utils.LogINFO(err.Error())
// 		return
// 	}

// 	res := response.ResponseStruct{
// 		HttpStatus: http.StatusOK,
// 		Code:       response.SuccessCode,
// 		Msg:        "ok",
// 		Data:       nil,
// 	}

// 	response.Response(c, res)
// }

// func GetSecondaryCategory(c *gin.Context) {
// 	var modelSecondaryCategories []model.SecondaryCategory
// 	common.DB.Find(&modelSecondaryCategories)

// 	utils.LogINFO(fmt.Sprint(modelSecondaryCategories))

// 	res := response.ResponseStruct{
// 		HttpStatus: http.StatusOK,
// 		Code:       response.SuccessCode,
// 		Msg:        "ok",
// 		Data: gin.H{
// 			"SecondaryCategories": modelSecondaryCategories,
// 		},
// 	}

// 	response.Response(c, res)
// }
