package dto

type CreateCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ModifyCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
