package service

import (
	"avitopvz/internal/models"
	"avitopvz/internal/storage"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Login(t *testing.T) {
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
			name:           "Login Error",
			expectedUser:   &validUser,
			expectedResult: "",
			needError:      true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().Login(mock.Anything, &validUser).Return(nil, errors.New("Login Error")).Times(1)
				return db
			},
		},
		{
			name:           "Invalid Password",
			expectedUser:   &models.User{Email: "testuser", Password: "wrongpassword"},
			expectedResult: "",
			needError:      true,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().Login(mock.Anything, &models.User{Email: "testuser", Password: "wrongpassword"}).Return(&models.User{Email: "testuser", Password: "$2a$10$YvrnsrnhiGmlnHtG6Ok19.uW5zwecDzZOu4xYRaqb9dDsdsjOYh.q", Role: "moderator"}, nil).Times(1)
				return db
			},
		},
		{
			name:           "Successful Login",
			expectedUser:   &validUser,
			expectedResult: "",
			needError:      false,
			mockSetup: func(db *storage.MockInterface) *storage.MockInterface {
				db.EXPECT().Login(mock.Anything, &validUser).Return(&models.User{Email: "testuser", Password: "$2a$10$YvrnsrnhiGmlnHtG6Ok19.uW5zwecDzZOu4xYRaqb9dDsdsjOYh.q", Role: "moderator"}, nil).Times(1)
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
			token, err := serv.Login(context.Background(), tt.expectedUser)

			// Assert

			if tt.needError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
