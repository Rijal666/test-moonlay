package repositories

import (
	"todo-list/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindTodos() ([]models.Todo, error)
	GetTodo(ID int) (models.Todo, error)
	CreateTodo(Todo models.Todo) (models.Todo, error)
	DeleteTodo(Todo models.Todo) (models.Todo, error)
	UpdateTodo(Todo models.Todo) (models.Todo, error)
	GetSublistByTodoID(todoID int) ([]models.SubList, error)
	GetAllLists(page, pageSize int, search string, preloadSublist bool) ([]models.Todo, int, error)
}

func RepositoryTodo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error

	return todos, err
}

func (r *repository) GetTodo(ID int) (models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, ID).Error

	return todo, err
}

func (r *repository) CreateTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error

	return todo, err
}

func (r *repository) DeleteTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Delete(&todo).Scan(&todo).Error

	return todo, err
}

func (r *repository) UpdateTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Save(&todo).Error

	return todo, err
}

func (r *repository) GetSublistByTodoID(todoID int) ([]models.SubList, error) {
	var sublists []models.SubList
	err := r.db.Where("todo_id = ?", todoID).Find(&sublists).Error

	return sublists, err
}

func (r *repository) GetAllLists(page, pageSize int, search string, preloadSublist bool) ([]models.Todo, int, error) {
	var todos []models.Todo
	query := r.db.Model(&models.Todo{})

	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalTodos int64
	err := query.Count(&totalTodos).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&todos).Error
	if err != nil {
		return nil, 0, err
	}

	if preloadSublist {
		for i := range todos {
			sublists, err := r.GetSublistByTodoID(todos[i].ID)
			if err != nil {
				return nil, 0, err
			}

			var sublistResponses []models.SubListResponse
			for _, sublist := range sublists {
				sublistResponse := models.SubListResponse{
					ID: sublist.ID, Title: sublist.Title, Description: sublist.Description, Files: sublist.Files,
				}
				sublistResponses = append(sublistResponses, sublistResponse)
			}

			todos[i].Sublist = sublistResponses
		}
	}

	return todos, int(totalTodos), nil
}
