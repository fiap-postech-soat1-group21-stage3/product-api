package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fiap-postech-soat1-group21/product-api/product-api/adapter/repository"
	"github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	createProduct = `
	INSERT INTO "product" (.+)
	VALUES (.+)
	ON CONFLICT DO NOTHING RETURNING "id"
`

	updateProduct = `UPDATE "product" SET (.+) WHERE (.+)`
)

var (
	defaultID = uuid.MustParse("b4dacf92-7000-4523-9fab-166212acc92d")

	ctx = context.Background()

	create = &entity.Product{
		Name:        "Burguer de planta",
		Description: "Burguer com carne vegetal",
		Category:    entity.Burguers,
		Price:       10,
	}

	update = &entity.Product{
		ID:          defaultID,
		Description: "Burguer com carne e maionese vegetal",
		Price:       12,
	}
)

func Test_CreateProduct(t *testing.T) {
	t.Run("when everything goes ok, should create a product register", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		mock.ExpectBegin()
		mock.
			ExpectQuery(createProduct).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(defaultID))
		mock.ExpectCommit()

		r := repository.New(dbGorm)
		result, err := r.Product.Create(ctx, create)

		assert.NoError(t, err)
		assert.Equal(t, result.ID, defaultID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("when db returns unmapped error, should propagate the error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()
		wantErr := errors.New("iamanerror")

		mock.ExpectBegin()
		mock.ExpectQuery(createProduct).WillReturnError(wantErr)
		mock.ExpectRollback()

		r := repository.New(dbGorm)
		result, err := r.Product.Create(ctx, create)
		assert.ErrorIs(t, err, wantErr)
		assert.Nil(t, result)
	})
}

func Test_UpdateProduct(t *testing.T) {
	t.Run(
		"when everything goes as expected, should update batch on db", func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
			defer db.Close()

			mock.ExpectBegin()
			mock.ExpectExec(updateProduct).WillReturnResult(sqlmock.NewResult(1, 1))

			mock.ExpectCommit()

			r := repository.New(dbGorm)
			err := r.Product.Update(ctx, update)
			assert.NoError(t, err)

			if err = mock.ExpectationsWereMet(); err != nil {
				assert.NoError(t, err)
			}
		})

	t.Run(
		"when db returns unmapped error, should propagate the error", func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
			defer db.Close()
			wantErr := errors.New("iamanerror")

			mock.ExpectBegin()

			mock.ExpectExec(updateProduct).WillReturnError(wantErr)

			mock.ExpectRollback()

			r := repository.New(dbGorm)
			err := r.Product.Update(ctx, update)
			assert.ErrorIs(t, err, wantErr)

			if err = mock.ExpectationsWereMet(); err != nil {
				assert.NoError(t, err)
			}
		})
}

// func TestFetchNotify(t *testing.T) {
// 	t.Run("when everything goes ok, should return a notification list register", func(t *testing.T) {
// 		db, mock, _ := sqlmock.New()
// 		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
// 		defer db.Close()

// 		wantRows := &entity.Customer{
// 			ID:    uuid.MustParse("24ecd959-96fb-4394-86e9-eb8f51878bf5"),
// 			Name:  "bia",
// 			CPF:   "222222222",
// 			Email: "b@email.com",
// 		}

// 		rows := sqlmock.NewRows([]string{"id", "name", "cpf", "email"}).
// 			AddRow(
// 				wantRows.ID,
// 				wantRows.Name,
// 				wantRows.CPF,
// 				wantRows.Email,
// 			)

// 		mock.
// 			ExpectQuery(fetchNotification).WillReturnRows(rows)

// 		r := repository.New(dbGorm)
// 		got, err := r.Customer.RetrieveCustomer(ctx, customer)

// 		assert.NoError(t, err)
// 		assert.NoError(t, mock.ExpectationsWereMet())
// 		assert.Equal(t, wantRows, got)
// 	})

// t.Run("when db returns unmapped error, should propagate the error", func(t *testing.T) {
// 	db, mock, _ := sqlmock.New()
// 	dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
// 	defer db.Close()
// 	wantErr := errors.New("iamanerror")

// 	mock.
// 		ExpectQuery(fetchNotification).WillReturnError(wantErr)

// 	r := repository.New(dbGorm)
// 	got, err := r.Notify.FetchNotify(ctx)

// 	assert.ErrorIs(t, err, wantErr)
// 	assert.Nil(t, got)
// })
//}
