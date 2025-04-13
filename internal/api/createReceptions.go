package api

import (
	"avitopvz/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) CreateReceptions(c *gin.Context) {
	var newReceptions models.Receptions

	if err := c.ShouldBindJSON(&newReceptions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receptions, err := h.service.CreateReceptions(c.Request.Context(), &newReceptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receptions)
}
