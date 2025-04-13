package api

import (
	"avitopvz/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetAllPVZ(c *gin.Context) {
	var ListInfo models.GetAllPVZRequest

	if err := c.ShouldBindQuery(&ListInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pvzs, err := h.service.GetAllPVZ(c.Request.Context(), ListInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pvzs)
}
