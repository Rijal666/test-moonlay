package routes

import (
	"todo-list/handlers"
	"todo-list/pkg/middleware"
	"todo-list/pkg/mysql"
	"todo-list/repositories"

	"github.com/labstack/echo/v4"
)

func TodoRoutes(e *echo.Group) {
	todoRepository := repositories.RepositoryTodo(mysql.ConnDB)
	h := handlers.HandlerTodo(todoRepository)

	e.GET("/todos", h.FindTodos)
	e.GET("/todo/:id", h.GetTodo)
	e.POST("/todo", middleware.UploadFile(h.CreateTodo))
	e.DELETE("/todo/:id", h.DeleteTodo)
	e.PUT("/todo/:id", middleware.UploadFile(h.UpdateTodo))
	e.GET("/all-lists", h.GetAllLists)
}
