package api

import (
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

func TestUserHandler_DeleteLastProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("успешное удаление товара", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		pvzID := "some-pvz-id"

		mockService.On("DeleteLastProduct", mock.Anything, pvzID).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/product/delete/"+pvzID, nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = gin.Params{
			{Key: "pvzId", Value: pvzID},
		}
		ctx.Request = req

		handler.DeleteLastProduct(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "\"Товар удален\"", w.Body.String())
	})

	t.Run("внутренняя ошибка сервиса", func(t *testing.T) {
		mockService := service.NewMockInterface(t)
		handler := &UserHandler{service: mockService}

		pvzID := "some-pvz-id"

		mockService.On("DeleteLastProduct", mock.Anything, pvzID).Return(errors.New("internal error"))

		req := httptest.NewRequest(http.MethodDelete, "/product/delete/"+pvzID, nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = gin.Params{
			{Key: "pvzId", Value: pvzID},
		}
		ctx.Request = req

		handler.DeleteLastProduct(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "internal error", resp["error"])
	})

}
