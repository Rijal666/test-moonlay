package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	resultdto "todo-list/dto/result"
	sublistdto "todo-list/dto/sublist"
	"todo-list/models"
	"todo-list/repositories"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type handlerSubList struct {
	SubListRepository repositories.SubListRepository
	TodoRepository    repositories.TodoRepository
}

func HandlerSubList(SubListRepository repositories.SubListRepository, TodoRepository repositories.TodoRepository) *handlerSubList {
	return &handlerSubList{SubListRepository, TodoRepository}
}

func (h *handlerSubList) FindSubLists(c echo.Context) error {
	sublists, err := h.SubListRepository.FindSubLists()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if len(sublists) > 0 {
		return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Data for all todos was successfully obtained", Data: convertResponseSubLists(sublists)})
	} else {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: "No record found"})
	}
}

func (h *handlerSubList) GetSubList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var sublist models.SubList
	sublist, err := h.SubListRepository.GetSubList(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Product data successfully obtained", Data: convertResponseSubList(sublist)})
}

func (h *handlerSubList) CreateSubList(c echo.Context) error {

	dataFile := c.Get("dataFiles").([]string)
	stringDataFiles := strings.Join(dataFile, ",")

	todo, _ := strconv.Atoi(c.FormValue("todo_id"))
	request := sublistdto.SublistRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Files:       stringDataFiles,
		TodoID:      todo,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})

	}

	todoid, err := h.TodoRepository.GetTodo(request.TodoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	sublist := models.SubList{
		Title:       request.Title,
		Description: request.Description,
		Files:       request.Files,
		TodoID:      request.TodoID,
		Todo:        convertTodoResponse(todoid),
	}
	data, err := h.SubListRepository.CreateSubList(sublist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	result := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"files":       strings.Split(data.Files, ","),
		"todo_id":     data.TodoID,
		"todo":        data.Todo,
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data sudah nambah", Data: result})

}

func (h *handlerSubList) UpdateSubList(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)
	todo, _ := strconv.Atoi(c.FormValue("todo_id"))
	request := sublistdto.UpdateSublistRequest{
		TodoID:      todo,
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Files:       dataFile,
	}

	id, _ := strconv.Atoi(c.Param("id"))
	sublist, err := h.SubListRepository.GetSubList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.TodoID != 0 {
		sublist.TodoID = request.TodoID
	}
	if request.Title != "" {
		sublist.Title = request.Title
	}
	if request.Description != "" {
		sublist.Description = request.Description
	}
	if request.Files != "" {
		sublist.Files = request.Files
	}

	data, err := h.SubListRepository.UpdateSubList(sublist)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data sudah berhasil di update", Data: convertResponseSubList(data)})

}

func (h *handlerSubList) DeleteSubList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sublist, err := h.SubListRepository.GetSubList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})

	}
	data, err := h.SubListRepository.DeleteSubList(sublist)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "data sudah berhasil di hapus", Data: convertResponseSubList(data)})
}

func (h *handlerSubList) GetAllSubLists(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	search := c.QueryParam("search")
	preloadSublist := c.QueryParam("preload_sublist") == "true"

	sublists, totalCount, err := h.SubListRepository.GetAllSubLists(page, pageSize, search, preloadSublist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	responseData := struct {
		TotalCount int              `json:"total_count"`
		Page       int              `json:"page"`
		PageSize   int              `json:"page_size"`
		Data       []models.SubList `json:"data"`
	}{
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		Data:       sublists,
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Todo data successfully obtained", Data: responseData})
}

func convertResponseSubList(u models.SubList) sublistdto.SubListResponse {
	return sublistdto.SubListResponse{
		ID:     u.ID,
		TodoID: u.TodoID,
		// Todo:        convertTodoResponse(),
		Title:       u.Title,
		Description: u.Description,
		Files:       u.Files,
	}
}

func convertResponseSubLists(sublists []models.SubList) []sublistdto.SubListResponse {
	var responseSubLists []sublistdto.SubListResponse

	for _, sublist := range sublists {
		responseSubLists = append(responseSubLists, convertResponseSubList(sublist))
	}

	return responseSubLists
}

func convertTodoResponse(c models.Todo) models.TodoResponse {
	return models.TodoResponse{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		Files:       c.Files,
	}
}
