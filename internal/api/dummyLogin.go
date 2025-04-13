package api

import (
	"avitopvz/internal/auth"
	"avitopvz/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) DummyLogin(c *gin.Context) {
	var testToken *models.User
	if err := c.ShouldBindJSON(&testToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	testToken.Email = "test"
	testToken.Password = "test"

	tokenString, err := auth.CreateToken(testToken, "secretkey")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при создании токена"})
		return
	}
	c.Header("Authorization", "Bearer "+tokenString)

	c.JSON(http.StatusOK, gin.H{"access_token": tokenString})
}
