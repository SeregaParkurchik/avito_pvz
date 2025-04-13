package api

import (
	"avitopvz/internal/models"
	"avitopvz/internal/service"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("успешный логин", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		user := models.User{
			Email:    "testuser",
			Password: "password",
		}

		tokenString := "test-token"

		mockService.On("Login", mock.Anything, &user).Return(tokenString, nil)

		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.Login(ctx)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, tokenString, resp["access_token"])
		assert.Equal(t, "Bearer "+tokenString, w.Header().Get("Authorization"))
	})

	t.Run("некорректные учетные данные", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		user := models.User{
			Email:    "testuser",
			Password: "wrongpassword",
		}

		mockService.On("Login", mock.Anything, &user).Return("", errors.New("invalid credentials"))

		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.Login(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "invalid credentials", resp["error"])
	})

	t.Run("некорректный JSON в запросе", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.Login(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp["error"], "invalid character")
	})
}
