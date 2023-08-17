package models

type SubList struct {
	ID          int          `json:"id" gorm:"primary_key:auto_increment"`
	TodoID      int          `json:"todo_id" gorm:"index"`
	Todo        TodoResponse `json:"todo"`
	Title       string       `json:"title" gorm:"type:varchar(100);not null"`
	Description string       `json:"description" gorm:"type:text;not null"`
	Files       string       `json:"files" gorm:"type:varchar(255)"`
}

type SubListResponse struct {
	ID          int    `json:"id"`
	TodoID      int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Files       string `json:"files"`
}

func (SubListResponse) TableName() string {
	return "sublists"
}
