package auth

import (
	"avitopvz/internal/models"
	"fmt"
	"os"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func Test_HashPassword(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name          string
		password      string
		needError     bool
		expectedError error
	}

	testCases := []testCase{
		{
			name:      "Successful HashPassword",
			password:  "password123",
			needError: false,
		},
		{
			name:      "Empty Password",
			password:  "",
			needError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			hashedPassword, err := HashPassword(tt.password)

			// Assert

			if tt.needError {
				require.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(tt.password))
				assert.NoError(t, err)
			}
		})
	}
}

func Test_GetJWTKey(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name          string
		setupEnv      func()
		needError     bool
		expectedError error
		expectedKey   string
	}

	testCases := []testCase{
		{
			name: "Empty JWT Key",
			setupEnv: func() {
				os.Setenv("JWT_SECRET_KEY", "")
			},
			needError:     true,
			expectedError: fmt.Errorf("нет секреного ключа для JWT"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			originalDir, err := os.Getwd()
			require.NoError(t, err)

			if tt.setupEnv != nil {
				tt.setupEnv()
			}

			// Act
			key, err := GetJWTKey()

			// Assert

			if tt.needError {
				require.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedKey, key)
			}

			// Cleanup
			os.Unsetenv("JWT_SECRET_KEY")
			os.Chdir(originalDir)
		})
	}
}

func Test_CreateToken(t *testing.T) {
	t.Parallel()

	validUser := &models.User{
		ID:    uuid.Must(uuid.NewV4()),
		Email: "test@example.com",
		Role:  "employee",
	}

	validKey := "test-secret-key"

	type testCase struct {
		name      string
		user      *models.User
		key       string
		needError bool
	}

	testCases := []testCase{
		{
			name:      "Successful CreateToken",
			user:      validUser,
			key:       validKey,
			needError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			tokenString, err := CreateToken(tt.user, tt.key)

			// Assert

			if tt.needError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					return []byte(tt.key), nil
				})
				require.NoError(t, err)

				claims, ok := token.Claims.(jwt.MapClaims)
				require.True(t, ok)

				assert.Equal(t, tt.user.ID.String(), claims["id"])
				assert.Equal(t, tt.user.Email, claims["email"])
				assert.Equal(t, tt.user.Role, claims["role"])
				assert.NotNil(t, claims["exp"])
			}
		})
	}
}
