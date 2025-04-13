package service

import (
	"avitopvz/internal/storage"
	"context"
	"errors"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_DeleteLastProduct(t *testing.T) {
	t.Parallel()

	validPvzID := uuid.Must(uuid.NewV4())
	validPvzIDStr := validPvzID.String()

	type testCase struct {
		name          string
		pvzIDStr      string
		needError     bool
		mockSetup     func(db *storage.MockInterface) *storage.MockInterface
		expectedError error
	}

	testCases := []testCase{
		{
			name:      "Invalid PVZ ID",
			pvzIDStr:  "invalid-uuid",
			needError: true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				return db
			},
			expectedError: errors.New("неверный ID ПВЗ"),
		},
		{
			name:      "DeleteLastProduct Error",
			pvzIDStr:  validPvzIDStr,
			needError: true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().DeleteLastProduct(mock.Anything, validPvzID).Return(errors.New("DeleteLastProduct Error")).Times(1)
				return db
			},
			expectedError: errors.New("DeleteLastProduct Error"),
		},
		{
			name:      "Successful DeleteLastProduct",
			pvzIDStr:  validPvzIDStr,
			needError: false,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().DeleteLastProduct(mock.Anything, validPvzID).Return(nil).Times(1)
				return db
			},
			expectedError: nil,
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
			err := serv.DeleteLastProduct(context.Background(), tt.pvzIDStr)

			// Assert

			if tt.needError {
				require.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
