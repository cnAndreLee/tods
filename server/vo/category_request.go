package vo

type CreateCategory struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"`
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
