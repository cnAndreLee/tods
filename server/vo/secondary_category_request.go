package vo

type CreateSecondaryCategory struct {
	ID       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	FatherID string `json:"fatherid" binding:"required"`
}
