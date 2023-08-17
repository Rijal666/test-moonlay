package routes

import (
	"todo-list/handlers"
	"todo-list/pkg/middleware"
	"todo-list/pkg/mysql"
	"todo-list/repositories"

	"github.com/labstack/echo/v4"
)

func SubListRoutes(e *echo.Group) {
	sublistRepository := repositories.RepositorySubList(mysql.ConnDB)
	todoRepository := repositories.RepositoryTodo(mysql.ConnDB)
	h := handlers.HandlerSubList(sublistRepository, todoRepository)

	e.GET("/sublists", h.FindSubLists)
	e.GET("/sublist/:id", h.GetSubList)
	e.POST("/sublist", middleware.UploadFile(h.CreateSubList))
	e.DELETE("/sublist/:id", h.DeleteSubList)
	e.PUT("/sublist/:id", middleware.UploadFile(h.UpdateSubList))
	e.GET("/all-sublists", h.GetAllSubLists)

}
