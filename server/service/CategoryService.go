package service

import (
	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/utils"
)

func IsExistChild(doCategory *model.Category) bool {
	if doCategory.Level == 1 {
		// 查找parentid为doCategory.ID的item
		result := common.DB.Where("parent_id = ?", doCategory.ID).First(&model.Category{})
		// 如果存在child，则返回true
		if result.Error == nil {
			utils.LogINFO("level==1, 存在child")
			return true
		}
	} else if doCategory.Level == 2 {
		// 在files表种查询是否有子文件
		result := common.DB.Where("file_belong = ?", doCategory.ID).First(&model.File{})
		if result.Error == nil {
			utils.LogINFO("level==2, 存在子文件")
			return true
		}
	}

	return false
}
