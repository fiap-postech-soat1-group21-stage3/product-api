package repository

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MaxOpenConns       = 2
	MaxIdleConns       = 2
	ConnMaxLifetimeSec = 100
	ConnMaxIdleSec     = 100
)

func NewRepository() (repo *Repository) {

	postgresDSN := os.Getenv("POSTGRES_DSN")

	gdb, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to create postgres db: %v", err)
	}

	repo = New(gdb)

	if MaxOpenConns > 0 {
		repo.SetMaxOpenConns(MaxOpenConns)
	}

	if MaxIdleConns > 0 {
		repo.SetMaxIdleConns(MaxIdleConns)
	}

	if ConnMaxLifetimeSec > 0 {
		repo.SetConnMaxLifetime(ConnMaxLifetimeSec)
	}

	if ConnMaxIdleSec > 0 {
		repo.SetConnMaxIdleTime(ConnMaxIdleSec)
	}

	return
}
