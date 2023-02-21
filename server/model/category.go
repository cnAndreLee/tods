package model

import (
	pq "github.com/lib/pq"
)

// DO:
// 一级分类parentID为0，LEVEL为1
// 二级分类parentID为父id，LEVEL为2
type Category struct {
	ID          string         `gorm:"type:varchar(50);primaryKey;" json:"id"`
	ParentID    string         `gorm:"type:varchar(50);not null;" json:"parent_id"`
	Level       uint8          `gorm:"type:smallint;not null" json:"level" `
	Title       string         `gorm:"type:varchar(100);not null;unique" json:"title"`
	Permissions pq.StringArray `gorm:"type:text[]" json:"permissions"`
}

// DTO: -----------------------------------
type CreateCategoryDTO struct { // !!!!!!!
	ParentID string `json:"parent_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Level    uint8  `json:"level" binding:"required"`
}

type DeleteCategoryDTO struct {
	ID string `json:"id" binding:"required"`
}

// 更名
type ModifyCategoryDTO struct {
	ID       string `json:"id" binding:"required"`
	NewTitle string `json:"new_title" binding:"required"`
}

// 修改权限
type ChangeCategoryPermissionDTO struct {
	ID             string         `json:"id" binding:"required"`
	NewPermissions pq.StringArray `json:"new_permissions" binding:"required"`
}

// VO: ---------------------------------------
type GetCategoryVO struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Level       uint8           `json:"level"`
	Children    []GetCategoryVO `json:"children"`
	Permissions pq.StringArray  `json:"permissions"`
}
