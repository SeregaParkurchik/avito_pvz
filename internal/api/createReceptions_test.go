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

func TestUserHandler_CreateReceptions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("успешное создание receptions", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		newReceptions := models.Receptions{
			PVZID:  uuid.Must(uuid.NewV4()),
			Status: "open",
		}

		createdReceptions := models.Receptions{
			ID:       uuid.Must(uuid.NewV4()),
			DateTime: time.Now(),
			PVZID:    uuid.Must(uuid.NewV4()),
			Status:   "open",
		}

		mockService.On("CreateReceptions", mock.Anything, &newReceptions).Return(&createdReceptions, nil)

		body, _ := json.Marshal(newReceptions)
		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.CreateReceptions(ctx)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp models.Receptions
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, createdReceptions.ID, resp.ID)
		assert.Equal(t, createdReceptions.PVZID, resp.PVZID)
		assert.Equal(t, createdReceptions.Status, resp.Status)
		assert.WithinDuration(t, createdReceptions.DateTime, resp.DateTime, time.Second)
	})

	t.Run("внутренняя ошибка сервиса", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		newReceptions := models.Receptions{
			PVZID:  uuid.Must(uuid.NewV4()),
			Status: "open",
		}

		mockService.On("CreateReceptions", mock.Anything, &newReceptions).Return(nil, errors.New("internal error"))

		body, _ := json.Marshal(newReceptions)
		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.CreateReceptions(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "internal error", resp["error"])
	})

	t.Run("некорректный JSON в запросе", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.CreateReceptions(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp["error"], "invalid character")
	})
}
