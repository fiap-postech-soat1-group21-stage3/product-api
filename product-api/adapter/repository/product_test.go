package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/adapter/repository"
	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/internal/domain/entity"
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

	selectProductList = `SELECT (.+) FROM "product"`

	selectCountProductList = `SELECT count(.+) FROM "product"`
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

func TestRetrieveProduct(t *testing.T) {
	t.Run("when everything goes ok, should return a product list", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()

		wantRows := &entity.ProductResponseList{
			Result: []*entity.Product{
				{
					ID:          uuid.MustParse("24ecd959-96fb-4394-86e9-eb8f51878bf5"),
					Name:        "hamburguer",
					Description: "hamburger com queijo",
					Category:    "hamburguer",
					Price:       10,
				},
			},
			Count: 1,
		}

		rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "price"}).
			AddRow(
				wantRows.Result[0].ID,
				wantRows.Result[0].Name,
				wantRows.Result[0].Description,
				wantRows.Result[0].Category,
				wantRows.Result[0].Price,
			)

		countRow := sqlmock.NewRows([]string{"count"}).AddRow(1)

		mock.
			ExpectQuery(selectProductList).WillReturnRows(rows)

		mock.
			ExpectQuery(selectCountProductList).
			WillReturnRows(countRow)

		r := repository.New(dbGorm)
		got, err := r.Product.GetProducts(ctx)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.Equal(t, wantRows.Result, got.Result)
	})

	t.Run("when db returns unmapped error, should propagate the error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		defer db.Close()
		wantErr := errors.New("iamanerror")

		mock.
			ExpectQuery(selectProductList).WillReturnError(wantErr)

		r := repository.New(dbGorm)
		got, err := r.Product.GetProducts(ctx)

		assert.ErrorIs(t, err, wantErr)
		assert.Nil(t, got)
	})
}
