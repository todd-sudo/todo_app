package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/todd-sudo/todo_app"
)

// signUp регистрация пользователей
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// signIn авторизация пользователей
func (h *Handler) signIn(c *gin.Context) {

}
