package service

import (
	"net/http"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/dto"
	"github.com/cnAndreLee/tods_server/response"
)

func CreateCategoryService(category dto.CreateCategory) response.ResponseStruct {

	var res response.ResponseStruct

	common.DB.Debug().Create(category)

	res = response.ResponseStruct{
		HttpStatus: http.StatusBadRequest,
		Code:       0,
		Data:       nil,
		Msg:        "注册成功",
	}

	return res

}
