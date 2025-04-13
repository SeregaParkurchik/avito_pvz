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

func TestUserHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("успешная регистрация", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		newUser := models.User{
			Email:    "newuser",
			Password: "password",
		}

		tokenString := "new-token"

		mockService.On("Register", mock.Anything, &newUser).Return(tokenString, nil)

		body, _ := json.Marshal(newUser)
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.Register(ctx)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, tokenString, resp["access_token"])
		assert.Equal(t, "Bearer "+tokenString, w.Header().Get("Authorization"))
	})

	t.Run("ошибка при регистрации", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		newUser := models.User{
			Email:    "newuser",
			Password: "password",
		}

		mockService.On("Register", mock.Anything, &newUser).Return("", errors.New("registration failed"))

		body, _ := json.Marshal(newUser)
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.Register(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "registration failed", resp["error"])
	})

	t.Run("некорректный JSON в запросе", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.Register(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp["error"], "invalid character")
	})
}
