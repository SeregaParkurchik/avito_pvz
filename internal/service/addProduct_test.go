package service

import (
	"avitopvz/internal/models"
	"avitopvz/internal/storage"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_AddProduct(t *testing.T) {
	t.Parallel()

	validProduct := &models.Product{
		Type:         "электроника",
		ReceptionsID: uuid.Must(uuid.NewV4()),
		PVZID:        uuid.Must(uuid.NewV4()),
	}

	type testCase struct {
		name            string
		product         *models.Product
		needError       bool
		mockSetup       func(db *storage.MockInterface) *storage.MockInterface
		expectedError   error
		expectedProduct *models.Product
	}

	testCases := []testCase{
		{
			name:      "Invalid Product Type",
			product:   &models.Product{Type: "мебель"},
			needError: true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				return db
			},
			expectedError: errors.New("такого типа товаров нет"),
		},
		{
			name:      "AddProduct Error",
			product:   validProduct,
			needError: true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().AddProduct(mock.Anything, mock.AnythingOfType("*models.Product")).Return(nil, errors.New("AddProduct Error")).Times(1)
				return db
			},
			expectedError: errors.New("AddProduct Error"),
		},
		{
			name:      "Successful AddProduct",
			product:   validProduct,
			needError: false,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				productWithID := &models.Product{
					ID:           uuid.Must(uuid.NewV4()),
					DateTime:     time.Now().UTC(),
					Type:         validProduct.Type,
					ReceptionsID: validProduct.ReceptionsID,
					PVZID:        validProduct.PVZID,
				}
				db.EXPECT().AddProduct(mock.Anything, mock.AnythingOfType("*models.Product")).Return(productWithID, nil).Times(1)
				return db
			},
			expectedError: nil,
			expectedProduct: &models.Product{
				ID:           uuid.Must(uuid.NewV4()),
				DateTime:     time.Now().UTC(),
				Type:         validProduct.Type,
				ReceptionsID: validProduct.ReceptionsID,
				PVZID:        validProduct.PVZID,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			dbMock := tt.mockSetup(storage.NewMockInterface(t))
			serv := &service{
				storage: dbMock,
			}

			// Act
			addedProduct, err := serv.AddProduct(context.Background(), tt.product)

			// Assert

			if tt.needError {
				require.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedProduct.Type, addedProduct.Type)
				assert.Equal(t, tt.expectedProduct.ReceptionsID, addedProduct.ReceptionsID)
				assert.Equal(t, tt.expectedProduct.PVZID, addedProduct.PVZID)
				assert.NotEqual(t, uuid.Nil, addedProduct.ID)
				assert.NotEqual(t, time.Time{}, addedProduct.DateTime)
			}
		})
	}
}
