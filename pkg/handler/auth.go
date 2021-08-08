package handler

import (
	mybank "github.com/SiriusPluge/my-bank-service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) registration(c *gin.Context) {
	var input mybank.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
