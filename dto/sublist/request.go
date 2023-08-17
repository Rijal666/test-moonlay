package sublistdto

import "todo-list/models"

type SublistRequest struct {
	TodoID      int         `json:"todo_id"`
	Todo        models.Todo `json:"todo"`
	Title       string      `json:"title" form:"title"  validate:"required,max=100,alphanum"`
	Description string      `json:"description" form:"description" validate:"required,max=1000"`
	Files       string      `json:"files" form:"files"`
}

type UpdateSublistRequest struct {
	TodoID      int         `json:"todo_id"`
	Todo        models.Todo `json:"todo"`
	Title       string      `json:"title" form:"title"`
	Description string      `json:"description" form:"description"`
	Files       string      `json:"files" form:"files"`
}
