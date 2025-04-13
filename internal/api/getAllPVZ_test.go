package api

import (
	"avitopvz/internal/models"
	"avitopvz/internal/service"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_GetAllPVZ(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		mockService.On("GetAllPVZ", mock.Anything, mock.Anything).Return([]models.PVZWithReceptions{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/pvz?startDate=2023-10-26T10:00:00Z&endDate=2023-10-27T10:00:00Z&page=1&limit=10", nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.GetAllPVZ(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		var actualPVZs []models.PVZWithReceptions
		err := json.Unmarshal(w.Body.Bytes(), &actualPVZs)
		assert.NoError(t, err)
	})

	t.Run("bad request - invalid query params", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		req := httptest.NewRequest(http.MethodGet, "/pvz?startDate=invalid&endDate=invalid&page=invalid&limit=invalid", nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.GetAllPVZ(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp["error"], "parsing time")
	})

	t.Run("internal server error", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		mockService.On("GetAllPVZ", mock.Anything, mock.Anything).Return(nil, errors.New("internal error"))

		req := httptest.NewRequest(http.MethodGet, "/pvz?startDate=2023-10-26T10:00:00Z&endDate=2023-10-27T10:00:00Z&page=1&limit=10", nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.GetAllPVZ(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "internal error", resp["error"])
	})
}
