package model

type User struct {
	Account string `gorm:"primaryKey;type:varchar(20)" json:"account"`
	Key     string `gorm:"type:varchar(50);not null" json:"key"`
	IsAdmin bool   `gorm:"type:boolean;not null" json:"is_admin"`
	School  string `gorm:"type:varchar(100)" json:"school"`
	OutDate string `gorm:"type:varchar(50); not null" json:"out_date"`
	Remark  string `gorm:"type:text" json:"remark"`
}

type DtoUserRegistry struct {
	Account string `json:"account" binding:"required"`
	Key     string `json:"key" binding:"required"`
	IsAdmin bool   `json:"is_admin"`
	School  string `json:"school"`
	OutDate string `json:"out_date" binding:"required"`
	Remark  string `json:"remark"`
}

type DtoUserLogin struct {
	Account string `json:"account" binding:"required"`
	Key     string `json:"key" binding:"required"`
}

type DtoChangeUserKey struct {
	Account string `json:"account" binding:"required"`
	NewKey  string `json:"new_key" binding:"required"`
}
