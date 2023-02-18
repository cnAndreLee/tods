package model

type File struct {
	FileID     string `gorm:"type:varchar(50);primaryKey;" json:"id"`
	FileBelong string `gorm:"type:varchar(50);not null" json:"file_belong"`
	Title      string `gorm:"type:varchar(100);not null;unique" json:"title"`
	Suffix     string `gorm:"type:varchar(10);not null" json:"suffix"`
}
