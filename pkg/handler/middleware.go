package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	AuthHeader = "Authorization"
	userCtxKey = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerSplited := strings.Split(bearer, " ")
	if len(headerSplited) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "not valid auth header")
		return
	}
	jwt := headerSplited[1]
	jwt = strings.Trim(jwt, " ")

	idUser, err := h.service.Authorization.CheckToken(jwt)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Set(userCtxKey, idUser)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtxKey)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return -1, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id wrong type")
		return -1, errors.New("user id wrong type")
	}

	return idInt, nil
}
