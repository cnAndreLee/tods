package model

// type Category struct {
// 	ID   string `gorm:"type:varchar(50);primaryKey;" json:"id"`
// 	Name string `gorm:"type:varchar(100);not null;unique" json:"name"`
// }

// type SecondaryCategory struct {
// 	ID       string `gorm:"type:varchar(50);primaryKey;" json:"id"`
// 	Name     string `gorm:"type:varchar(100);not null" json:"name"`
// 	FatherID string `gorm:"type:varchar(50);not null" json:"fatherid"`
// }

// DO:
// 一级分类parentID为0，LEVEL为1
// 二级分类parentID为父id，LEVEL为2
type Category struct {
	ID       string `gorm:"type:varchar(50);primaryKey;" json:"id"`
	ParentID string `gorm:"type:varchar(50);not null;" json:"parent_id"`
	Level    uint8  `gorm:"type:smallint;not null" json:"level" `
	Title    string `gorm:"type:varchar(100);not null;unique" json:"title"`
}

// DTO:
type CreateCategoryVO struct { // !!!!!!!
	ParentID string `json:"parent_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Level    uint8  `json:"level" binding:"required"`
}

type DeleteCategoryDTO struct {
	ID string `json:"id" binding:"required"`
}

type ModifyCategoryDTO struct {
	ID       string `json:"id" binding:"required"`
	NewTitle string `json:"new_title" binding:"required"`
}

// VO:
type GetCategoryVO struct {
	ID       string          `json:"id"`
	Title    string          `json:"title"`
	Children []GetCategoryVO `json:"children"`
}
