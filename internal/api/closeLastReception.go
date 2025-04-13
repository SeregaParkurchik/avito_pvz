package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) CloseLastReception(c *gin.Context) {
	pvzIDStr := c.Param("pvzId")

	reception, err := h.service.CloseLastReception(c.Request.Context(), pvzIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reception)
}
