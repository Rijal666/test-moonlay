package sublistdto

import "todo-list/models"

type SubListResponse struct {
	ID          int                 `json:"id"`
	TodoID      int                 `json:"todo_id"`
	Todo        models.TodoResponse `json:"todo"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Files       string              `json:"files"`
}
