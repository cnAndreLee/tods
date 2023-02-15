package vo

type CreateCategory struct {
	ID    string `json:"id"`
	Title string `json:"title" binding:"required"`
	Level uint8  `json:"level"`
}

type DeleteCategory struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name"`
}

type ModifyCategory struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"`
}

type GetCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
