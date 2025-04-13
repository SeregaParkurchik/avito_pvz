package api

import (
	"avitopvz/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) CreatePVZ(c *gin.Context) {
	var newPVZ models.PVZ

	if err := c.ShouldBindJSON(&newPVZ); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pvz, err := h.service.CreatePVZ(c.Request.Context(), &newPVZ)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pvz)
}
