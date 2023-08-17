package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	resultdto "todo-list/dto/result"
	tododto "todo-list/dto/todo"
	"todo-list/models"
	"todo-list/repositories"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type handlerTodo struct {
	TodoRepository repositories.TodoRepository
}

func HandlerTodo(TodoRepository repositories.TodoRepository) *handlerTodo {
	return &handlerTodo{TodoRepository}
}

func (h *handlerTodo) FindTodos(c echo.Context) error {
	todos, err := h.TodoRepository.FindTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if len(todos) > 0 {
		return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Data for all todos was successfully obtained", Data: todos})
	} else {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: "No record found"})
	}
}

func (h *handlerTodo) GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var todo models.Todo
	todo, err := h.TodoRepository.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	sublists, err := h.TodoRepository.GetSublistByTodoID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	var todoSublistResponses []models.SubListResponse
	for _, sublist := range sublists {
		sublistResponse := models.SubListResponse{
			ID: sublist.ID, Title: sublist.Title, Description: sublist.Description, Files: sublist.Files,
		}
		todoSublistResponses = append(todoSublistResponses, sublistResponse)
	}

	todo.Sublist = todoSublistResponses

	if todo.Sublist != nil {
		return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Product data successfully obtained", Data: todo})
	}

	todo.Sublist = make([]models.SubListResponse, 0)
	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Product data successfully obtained", Data: todo})
}

func (h *handlerTodo) CreateTodo(c echo.Context) error {

	dataFile := c.Get("dataFiles").([]string)
	stringDataFiles := strings.Join(dataFile, ",")
	request := tododto.TodoRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Files:       stringDataFiles,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})

	}

	todo := models.Todo{
		Title:       request.Title,
		Description: request.Description,
		Files:       request.Files,
	}
	data, err := h.TodoRepository.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	fmt.Println(data.Files, "dataFiles")
	result := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"files":       strings.Split(data.Files, ","),
		"sublist":     data.Sublist,
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data sudah nambah", Data: result})

}

func (h *handlerTodo) UpdateTodo(c echo.Context) error {
	dataFile := c.Get("dataFiles").([]string)
	stringDataFiles := strings.Join(dataFile, ",")

	request := tododto.UpdateTodoRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Files:       stringDataFiles,
	}

	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.TodoRepository.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		todo.Title = request.Title
	}
	if request.Description != "" {
		todo.Description = request.Description
	}
	if request.Files != "" {
		todo.Files = request.Files
	}

	data, err := h.TodoRepository.UpdateTodo(todo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})

	}
	result := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"files":       strings.Split(data.Files, ","),
		"sublist":     data.Sublist,
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data sudah berhasil di update", Data: result})

}

func (h *handlerTodo) DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.TodoRepository.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})

	}
	data, err := h.TodoRepository.DeleteTodo(todo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data sudah berhasil di hapus", Data: data})
}

func (h *handlerTodo) GetAllLists(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	search := c.QueryParam("search")
	preloadSublist := c.QueryParam("preload_sublist") == "true"

	todos, totalCount, err := h.TodoRepository.GetAllLists(page, pageSize, search, preloadSublist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	responseData := struct {
		TotalCount int           `json:"total_count"`
		Page       int           `json:"page"`
		PageSize   int           `json:"page_size"`
		Data       []models.Todo `json:"data"`
	}{
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		Data:       todos,
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Todo data successfully obtained", Data: responseData})
}
