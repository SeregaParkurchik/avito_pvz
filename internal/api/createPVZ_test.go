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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_CreatePVZ(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("успешное создание PVZ", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		newPVZ := models.PVZ{
			City: "Moscow",
		}

		createdPVZ := models.PVZ{
			ID:               uuid.Must(uuid.NewV4()),
			RegistrationDate: time.Now(),
			City:             "Moscow",
		}

		mockService.On("CreatePVZ", mock.Anything, &newPVZ).Return(&createdPVZ, nil)

		body, _ := json.Marshal(newPVZ)
		req := httptest.NewRequest(http.MethodPost, "/pvz", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.CreatePVZ(ctx)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp models.PVZ
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, createdPVZ.ID, resp.ID)
		assert.Equal(t, createdPVZ.City, resp.City)
		assert.WithinDuration(t, createdPVZ.RegistrationDate, resp.RegistrationDate, time.Second)
	})

	t.Run("внутренняя ошибка сервиса", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		newPVZ := models.PVZ{
			City: "Moscow",
		}

		mockService.On("CreatePVZ", mock.Anything, &newPVZ).Return(nil, errors.New("internal error"))

		body, _ := json.Marshal(newPVZ)
		req := httptest.NewRequest(http.MethodPost, "/pvz", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.CreatePVZ(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "internal error", resp["error"])
	})

	t.Run("некорректный JSON в запросе", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		req := httptest.NewRequest(http.MethodPost, "/pvz", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.CreatePVZ(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp["error"], "invalid character")
	})
}
