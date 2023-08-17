package tododto

type TodoRequest struct {
	Title       string `json:"title" form:"title"  validate:"required,max=100,alphanum"`
	Description string `json:"description" form:"description" validate:"required,max=1000"`
	Files       string `json:"files" form:"files"`
	// Sublist     []models.SubListResponse `json:"sublist"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Files       string `json:"files" form:"files"`
}
