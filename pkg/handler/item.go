package handler

import (
	"github.com/gin-gonic/gin"
	todo "go-todo"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.CreateTodoInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect list id")
		return
	}

	id, err := h.service.TodoItem.CreateItem(userId, listId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})

}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect list id")
		return
	}

	items, err := h.service.TodoItem.GetAllItems(userId, listId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)

}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect list id")
		return
	}
	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect item id")
		return
	}

	item, err := h.service.TodoItem.GetItemById(userId, listId, itemId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect list id")
		return
	}
	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect item id")
		return
	}

	err = h.service.TodoItem.DeleteItem(userId, listId, itemId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
