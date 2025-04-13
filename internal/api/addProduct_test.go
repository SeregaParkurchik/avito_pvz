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

func TestUserHandler_AddProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("успешное добавление продукта", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		productID := uuid.Must(uuid.NewV4())
		receptionID := uuid.Must(uuid.NewV4())
		now := time.Now()

		product := &models.Product{
			ID:           productID,
			DateTime:     now,
			Type:         "test-type",
			ReceptionsID: receptionID,
		}

		mockService.On("AddProduct", mock.Anything, mock.Anything).Return(product, nil)

		body, _ := json.Marshal(product)
		req := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.AddProduct(ctx)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp struct {
			ID           uuid.UUID `json:"id"`
			DateTime     time.Time `json:"dateTime"`
			Type         string    `json:"type"`
			ReceptionsID uuid.UUID `json:"receptionId"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, product.ID, resp.ID)
		assert.Equal(t, product.DateTime.Format(time.RFC3339), resp.DateTime.Format(time.RFC3339))
		assert.Equal(t, product.Type, resp.Type)
		assert.Equal(t, product.ReceptionsID, resp.ReceptionsID)

	})

	t.Run("ошибка при некорректном JSON", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		body := []byte(`{invalid-json}`)

		req := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.AddProduct(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockService.AssertNotCalled(t, "AddProduct", mock.Anything, mock.Anything)
	})

	t.Run("внутренняя ошибка сервиса", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		product := &models.Product{
			ID:           uuid.Must(uuid.NewV4()),
			DateTime:     time.Now(),
			Type:         "test-type",
			ReceptionsID: uuid.Must(uuid.NewV4()),
		}

		mockService.On("AddProduct", mock.Anything, mock.Anything).Return(nil, errors.New("something went wrong"))

		body, _ := json.Marshal(product)
		req := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		handler.AddProduct(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
