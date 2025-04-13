package service

import (
	"avitopvz/internal/models"
	"avitopvz/internal/storage"
	"context"
	"errors"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Register(t *testing.T) {
	t.Parallel()

	validUser := models.User{
		Email:    "testuser",
		Password: "123456",
		Role:     "moderator",
	}

	type testCase struct {
		name           string
		expectedUser   *models.User
		expectedResult string
		needError      bool
		mockSetup      func(db *storage.MockInterface) *storage.MockInterface
	}

	testCases := []testCase{
		{
			name:           "Invalid Role",
			expectedUser:   &models.User{Email: "test", Role: ""},
			expectedResult: "",
			needError:      true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {

				return db
			},
		},
		{
			name:           "Invalid Hash Password",
			expectedUser:   &models.User{Email: "test", Password: ""},
			expectedResult: "",
			needError:      true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				return db
			},
		},
		{
			name:           "RegisterError",
			expectedUser:   &validUser,
			expectedResult: "",
			needError:      true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().Register(mock.Anything, mock.MatchedBy(func(user *models.User) bool {
					return user.Email == "testuser" && user.Role == "moderator" && user.Password != "123456" && user.ID != uuid.Nil
				})).Return(errors.New("Register Error")).Times(1)

				return db
			},
		},
		{
			name:           "Success Register",
			expectedUser:   &validUser,
			expectedResult: "",
			needError:      false,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().Register(mock.Anything, mock.MatchedBy(func(user *models.User) bool {
					return user.Email == "testuser" && user.Role == "moderator" && user.Password != "123456" && user.ID != uuid.Nil
				})).Return(nil).Times(1)

				return db
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
			token, err := serv.Register(context.Background(), tt.expectedUser)

			// Assert

			if tt.needError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, token)
			}
			//assert.Equal(t, tt.expectedResult, token)

		})
	}
}
