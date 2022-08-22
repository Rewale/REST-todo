package handler

import (
	"github.com/gin-gonic/gin"
	todo "go-todo"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	id, err := h.service.TodoList.CreateList(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.service.TodoList.GetAllLists(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, getAllListsResponse{
		Data: lists,
	})

}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	list, err := h.service.TodoList.GetListById(userId, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, list)
}

func (h *Handler) updateList(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = h.service.TodoList.UpdateListById(userId, listId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	err = h.service.TodoList.DeleteListById(userId, listId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
