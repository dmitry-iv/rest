package sqlite

import (
	"errors"
	"fmt"
	"goland/cmd/iternal/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	Alias string `gorm:"unique;not null"`
	URL   string `gorm:"not null"`
}

type Storage struct {
	db *gorm.DB
}

func (s *Storage) GetAllURLs() ([]URL, error) {
	var urls []URL
	err := s.db.Find(&urls).Error
	return urls, err
}

//go:generate mockery --name=Storage --output=./mocks --outpkg=mocks --with-expecter
func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	dsn := "host=localhost user=postgres password=123 dbname=storage port=5432 sslmode=disable"
	// Использование:
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Миграции
	err = db.AutoMigrate(&URL{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.sqlite.SaveURL"

	url := URL{
		Alias: alias,
		URL:   urlToSave,
	}

	result := s.db.Create(&url)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrURLExist)
		}
		return 0, fmt.Errorf("%s: %w", op, result.Error)
	}

	return int64(url.ID), nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.sqlite.GetURL"

	var url URL
	result := s.db.Where("alias = ?", alias).First(&url)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", storage.ErrURLNotFound
	}
	if result.Error != nil {
		return "", fmt.Errorf("%s: %w", op, result.Error)
	}

	return url.URL, nil
}

//func (s *Storage) DeleteURL(alias string) error {
//const op = "storage.sqlite.DeleteURL"
//_, err := s.db.Prepare("DELETE FROM url WHERE alias = ?")
//if err != nil {
//return fmt.Errorf("%s: %w", op, err)
//}
//return nil
//}
