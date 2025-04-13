package api

import (
	"avitopvz/internal/models"
	"avitopvz/internal/service"
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

func TestUserHandler_CloseLastReception(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("успешное закрытие последней приёмки", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		pvzID := "some-pvz-id"
		reception := &models.Receptions{
			ID:       uuid.Must(uuid.NewV4()),
			DateTime: time.Now(),
			PVZID:    uuid.Must(uuid.NewV4()),
			Status:   "close",
		}

		mockService.On("CloseLastReception", mock.Anything, pvzID).Return(reception, nil)

		req := httptest.NewRequest(http.MethodPost, "/reception/close/"+pvzID, nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = gin.Params{
			{Key: "pvzId", Value: pvzID},
		}
		ctx.Request = req

		handler.CloseLastReception(ctx)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp models.Receptions
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, reception.ID, resp.ID)
		assert.Equal(t, reception.PVZID, resp.PVZID)
		assert.Equal(t, reception.Status, resp.Status)
		assert.WithinDuration(t, reception.DateTime, resp.DateTime, time.Second)
	})

	t.Run("внутренняя ошибка сервиса", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		pvzID := "some-pvz-id"

		mockService.On("CloseLastReception", mock.Anything, pvzID).Return(nil, errors.New("internal error"))

		req := httptest.NewRequest(http.MethodPost, "/reception/close/"+pvzID, nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = gin.Params{
			{Key: "pvzId", Value: pvzID},
		}
		ctx.Request = req

		handler.CloseLastReception(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var body map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &body)
		assert.NoError(t, err)
		assert.Equal(t, "internal error", body["error"])
	})
}
