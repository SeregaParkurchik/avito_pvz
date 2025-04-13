package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) DeleteLastProduct(c *gin.Context) {
	pvzIDStr := c.Param("pvzId")

	err := h.service.DeleteLastProduct(c.Request.Context(), pvzIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Товар удален")

}
