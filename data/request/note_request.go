package request

type CreateNoteRequest struct {
	Content string `json:"content" validate:"required,min=2,max=100"`
}

type UpdateNoteRequest struct {
	Id      int    `validate:"required"`
	Content string `validate:"required,max=200,min=2" json:"content"`
}
