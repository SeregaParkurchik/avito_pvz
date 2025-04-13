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

func Test_CreateReceptions(t *testing.T) {
	t.Parallel()

	validPvzID := uuid.Must(uuid.NewV4())

	type testCase struct {
		name               string
		newReceptions      *models.Receptions
		needError          bool
		mockSetup          func(db *storage.MockInterface) *storage.MockInterface
		expectedError      error
		expectedReceptions *models.Receptions
	}

	testCases := []testCase{
		{
			name: "Storage Error",
			newReceptions: &models.Receptions{
				PVZID: validPvzID,
			},
			needError: true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().CreateReceptions(mock.Anything, mock.AnythingOfType("*models.Receptions")).Return(nil, errors.New("Storage Error")).Times(1)
				return db
			},
			expectedError: errors.New("Storage Error"),
		},
		{
			name: "Success",
			newReceptions: &models.Receptions{
				PVZID: validPvzID,
			},
			needError: false,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				expectedReceptions := &models.Receptions{
					ID:       uuid.Must(uuid.NewV4()),
					PVZID:    validPvzID,
					Status:   "in_progress",
					DateTime: time.Now().UTC(),
				}
				db.EXPECT().CreateReceptions(mock.Anything, mock.AnythingOfType("*models.Receptions")).Return(expectedReceptions, nil).Times(1)
				return db
			},
			expectedReceptions: &models.Receptions{
				PVZID:    validPvzID,
				Status:   "in_progress",
				DateTime: time.Now().UTC(),
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
			receptions, err := serv.CreateReceptions(context.Background(), tt.newReceptions)

			// Assert

			if tt.needError {
				require.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedReceptions.PVZID, receptions.PVZID)
				assert.Equal(t, tt.expectedReceptions.Status, receptions.Status)
				assert.WithinDuration(t, tt.expectedReceptions.DateTime, receptions.DateTime, time.Second)
				assert.NotEqual(t, uuid.Nil, receptions.ID)
			}
		})
	}
}
