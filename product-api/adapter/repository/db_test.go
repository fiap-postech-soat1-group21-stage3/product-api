package repository_test

import (
	"os"
	"testing"

	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/adapter/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetPostgresURL(t *testing.T) {
	tests := []struct {
		name string
		dsn  string
	}{

		{
			name: "when url is invalid; should panic",
			dsn:  "%",
		},
		{
			name: "when url is empty; should panic",
			dsn:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("POSTGRES_DSN", tt.dsn)
			assert.Panics(t, func() { repository.NewRepository() })
		})
	}
}
