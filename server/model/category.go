package model

type Category struct {
	ID   string `gorm:"type:varchar(50);primaryKey;" json:"id"`
	Name string `gorm:"type:varchar(100);not null;unique" json:"name"`
}

type SecondaryCategory struct {
	ID       string `gorm:"type:varchar(50);primaryKey;" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	FatherID string `gorm:"type:varchar(50);not null" json:"fatherid"`
}
