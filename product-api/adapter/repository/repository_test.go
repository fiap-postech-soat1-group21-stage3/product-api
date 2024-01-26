package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fiap-postech-soat1-group21/product-api/product-api/adapter/repository"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	t.Run(
		"when everything goes ok, should return no error", func(t *testing.T) {
			db, _, _ := sqlmock.New()
			dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
			defer db.Close()

			r := repository.New(dbGorm)
			assert.NotNil(t, r)
		})
}

func TestRepository_SetMaxOpenConns(t *testing.T) {
	t.Run("when everything goes ok, should return no error", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		r := repository.New(dbGorm)

		r.SetMaxOpenConns(3)

		db, _ = r.Conn.DB()
		value := db.Stats().MaxOpenConnections
		assert.Equal(t, value, 3)
	})
}

func TestRepository_SetMaxIdleConns(t *testing.T) {
	t.Run("when everything goes ok, should return no error", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		r := repository.New(dbGorm)

		r.SetMaxIdleConns(3)
	})
}

func TestRepository_SetConnMaxLifetime(t *testing.T) {
	t.Run("when everything goes ok, should return no error", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		r := repository.New(dbGorm)

		r.SetConnMaxLifetime(time.Hour)
	})
}

func TestRepository_SetConnMaxIdleTime(t *testing.T) {
	t.Run("when everything goes ok, should return no error", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		r := repository.New(dbGorm)

		r.SetConnMaxIdleTime(time.Hour)
	})
}
