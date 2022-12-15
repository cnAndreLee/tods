package model

type User struct {
	// gorm.Model
	Account    string `gorm:"primaryKey;type:varchar(20)" json:"account"`
	Class      string `gorm:"type:varchar(20)" json:"class"`
	Key        string `gorm:"type:varchar(50);not null" json:"key"`
	Schoolname string `gorm:"type:varchar(100)" json:"schoolname"`
	Telephone  string `gorm:"type:varchar(20)" json:"telephone"`
	Outtime    string `gorm:"type:varchar(50)" json:"outtime"`
}
