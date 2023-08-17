package models

type Todo struct {
	ID          int               `json:"id" gorm:"primary_key:auto_increment"`
	Title       string            `json:"title" gorm:"type:varchar(100)"`
	Description string            `json:"description" gorm:"type:text"`
	Files       string            `json:"files" gorm:"type:varchar(255)"`
	Sublist     []SubListResponse `json:"sublist"`
}

type TodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Files       string `json:"files"`
}

func (TodoResponse) TableName() string {
	return "todos"
}
