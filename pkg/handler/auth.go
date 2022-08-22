package handler

import (
	"github.com/gin-gonic/gin"
	todo "go-todo"
	"net/http"
	"strings"
)

func (h *Handler) singUp(c *gin.Context) {
	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, map[string]any{
		"id": id,
	})

}

type singInInput struct {
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binging:"required"`
}

func (h *Handler) singIn(c *gin.Context) {
	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	jwt, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, map[string]any{
		"token": jwt,
	})

}
func (h *Handler) CheckJWT(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	jwt := strings.Split(bearer, " ")[1]
	jwt = strings.Trim(jwt, " ")

	username, err := h.service.Authorization.CheckToken(jwt)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, map[string]any{
		"id_user": username,
	})

}
