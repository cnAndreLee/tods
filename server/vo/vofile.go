package vo

type UploadFile struct {
	Name string `json:"name" binding:"required"`
}
