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

func Test_CloseLastReception(t *testing.T) {
	t.Parallel()

	validPvzID := uuid.Must(uuid.NewV4())
	validPvzIDStr := validPvzID.String()

	type testCase struct {
		name               string
		pvzIDStr           string
		needError          bool
		mockSetup          func(db *storage.MockInterface) *storage.MockInterface
		expectedError      error
		expectedReceptions *models.Receptions
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
			name:      "CloseLastReception Error",
			pvzIDStr:  validPvzIDStr,
			needError: true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().CloseLastReception(mock.Anything, validPvzID).Return(nil, errors.New("CloseLastReception Error")).Times(1)
				return db
			},
			expectedError: errors.New("CloseLastReception Error"),
		},
		{
			name:     "Success",
			pvzIDStr: validPvzIDStr,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				expectedReceptions := &models.Receptions{
					ID:       uuid.Must(uuid.NewV4()),
					PVZID:    validPvzID,
					Status:   "closed",
					DateTime: time.Now().UTC(),
				}
				db.EXPECT().CloseLastReception(mock.Anything, validPvzID).Return(expectedReceptions, nil).Times(1)
				return db
			},
			expectedReceptions: &models.Receptions{
				ID:       uuid.Must(uuid.NewV4()),
				PVZID:    validPvzID,
				Status:   "closed",
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
			receptions, err := serv.CloseLastReception(context.Background(), tt.pvzIDStr)

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
