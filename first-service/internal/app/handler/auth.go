package handler

import (
	"gin_prometheus/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.Services.Auth.CreateUser(input); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *Handler) SaySomething(c *gin.Context) {
	_, err := c.Writer.Write([]byte("Somebody says something:)"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

}
