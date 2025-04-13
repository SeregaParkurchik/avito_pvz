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

func Test_CreatePVZ(t *testing.T) {
	t.Parallel()

	validPVZ := &models.PVZ{
		ID:               uuid.Must(uuid.NewV4()),
		City:             "Москва",
		RegistrationDate: time.Now(),
	}

	type testCase struct {
		name          string
		newPVZ        *models.PVZ
		needError     bool
		mockSetup     func(db *storage.MockInterface) *storage.MockInterface
		expectedError error
	}

	testCases := []testCase{
		{
			name:      "Invalid City",
			newPVZ:    &models.PVZ{City: "Владивосток"},
			needError: true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				return db
			},
			expectedError: errors.New("в данном городе нет пвз, либо город указан неверно"),
		},
		{
			name:      "CreatePVZ Error",
			newPVZ:    validPVZ,
			needError: true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().CreatePVZ(mock.Anything, validPVZ).Return(nil, errors.New("CreatePVZ Error")).Times(1)
				return db
			},
			expectedError: errors.New("CreatePVZ Error"),
		},
		{
			name:      "Successful CreatePVZ",
			newPVZ:    validPVZ,
			needError: false,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().CreatePVZ(mock.Anything, validPVZ).Return(validPVZ, nil).Times(1)
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
			createdPVZ, err := serv.CreatePVZ(context.Background(), tt.newPVZ)

			// Assert

			if tt.needError {
				require.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.newPVZ.ID, createdPVZ.ID)
				assert.Equal(t, tt.newPVZ.City, createdPVZ.City)
				assert.Equal(t, tt.newPVZ.RegistrationDate.UTC(), createdPVZ.RegistrationDate)
			}
		})
	}
}
