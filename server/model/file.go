package model

type File struct {
	FileID     string `gorm:"type:varchar(50);primaryKey;" json:"id"`
	FileBelong string `gorm:"type:varchar(50);not null" json:"filebelong"`
	Title      string `gorm:"type:varchar(100);not null" json:"title"`
	Suffix     string `gorm:"type:varchar(10);not null" json:"suffix"`
}

type ResFile struct {
	FileID     string `json:"id"`
	FileBelong string `json:"filebelong"`
	Title      string `json:"title"`
	Suffix     string `json:"suffix"`
	Url        string `json:"url"`
}
