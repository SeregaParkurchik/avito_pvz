package api

import (
	"avitopvz/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func (h *UserHandler) AddProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addProduct, err := h.service.AddProduct(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := struct {
		ID           uuid.UUID `json:"id"`
		DateTime     time.Time `json:"dateTime"`
		Type         string    `json:"type"`
		ReceptionsID uuid.UUID `json:"receptionId"`
	}{
		ID:           addProduct.ID,
		DateTime:     addProduct.DateTime,
		Type:         addProduct.Type,
		ReceptionsID: addProduct.ReceptionsID,
	}

	c.JSON(http.StatusOK, response)
}
