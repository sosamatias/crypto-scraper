package repository

import (
	"github.com/sosamatias/crypto-scraper/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository interface {
	CreateInBatches(cryptoSnapshots []model.CryptoSnapshot, batchSize int) error
	Count() (int64, error)
}

func NewRepository(dbFile string) (Repository, error) {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&model.CryptoSnapshot{},
	)
	return repository{
		DB: db,
	}, nil
}

type repository struct {
	DB *gorm.DB
}

func (r repository) CreateInBatches(cryptoSnapshots []model.CryptoSnapshot, batchSize int) error {
	tx := r.DB.CreateInBatches(cryptoSnapshots, batchSize)
	return tx.Error
}

func (r repository) Count() (int64, error) {
	var count int64
	tx := r.DB.Model(&model.CryptoSnapshot{}).Count(&count)
	return count, tx.Error
}
