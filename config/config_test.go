package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		os.Setenv(httpHostEnvName, "localhost")
		os.Setenv(httpPortEnvName, "8080")
		defer os.Unsetenv(httpHostEnvName)
		defer os.Unsetenv(httpPortEnvName)

		cfg, err := NewHTTPConfig()
		assert.NoError(t, err)
		assert.Equal(t, "localhost:8080", cfg.Address())
	})

	t.Run("missing host", func(t *testing.T) {
		os.Unsetenv(httpHostEnvName)
		os.Setenv(httpPortEnvName, "8080")
		defer os.Unsetenv(httpPortEnvName)

		cfg, err := NewHTTPConfig()
		assert.Error(t, err)
		assert.Nil(t, cfg)
		assert.Equal(t, "http host not found", err.Error())
	})

	t.Run("missing port", func(t *testing.T) {
		os.Setenv(httpHostEnvName, "localhost")
		os.Unsetenv(httpPortEnvName)
		defer os.Unsetenv(httpHostEnvName)

		cfg, err := NewHTTPConfig()
		assert.Error(t, err)
		assert.Nil(t, cfg)
		assert.Equal(t, "http port not found", err.Error())
	})
}

func TestNewPGConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		os.Setenv(dsnEnvName, "postgres://user:password@host:port/dbname")
		defer os.Unsetenv(dsnEnvName)

		cfg, err := NewPGConfig()
		assert.NoError(t, err)
		assert.Equal(t, "postgres://user:password@host:port/dbname", cfg.DSN())
	})

	t.Run("missing dsn", func(t *testing.T) {
		os.Unsetenv(dsnEnvName)

		cfg, err := NewPGConfig()
		assert.Error(t, err)
		assert.Nil(t, cfg)
		assert.Equal(t, "pg dsn not found", err.Error())
	})
}

func TestTestPGConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		os.Setenv(dsnTestName, "postgres://testuser:testpassword@testhost:testport/testdbname")
		defer os.Unsetenv(dsnTestName)

		cfg, err := TestPGConfig()
		assert.NoError(t, err)
		assert.Equal(t, "postgres://testuser:testpassword@testhost:testport/testdbname", cfg.DSN())
	})

	t.Run("missing test dsn", func(t *testing.T) {
		os.Unsetenv(dsnTestName)

		cfg, err := TestPGConfig()
		assert.Error(t, err)
		assert.Nil(t, cfg)
		assert.Equal(t, "pg dsn not found", err.Error())
	})
}
