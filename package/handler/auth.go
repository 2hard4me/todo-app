package handler

import (
	"net/http"

	"github.com/2hard4me/todo-app"
	"github.com/gin-gonic/gin"
)



func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorRespone(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorRespone(c, http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})

}

func (h *Handler) signIn(c *gin.Context) {

}