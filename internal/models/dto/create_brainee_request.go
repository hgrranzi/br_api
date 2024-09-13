package dto

type CreateBraineeRequest struct {
	Text   string `json:"text" validate:"required,min=1,max=256"`
	Author string `json:"author" validate:"required,min=1,max=32"`
	Brand  string `json:"brand" validate:"required,min=1,max=32"`
}
