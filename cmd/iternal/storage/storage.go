package storage

import "errors"

// Storage определяет контракт для работы с хранилищем URL
type Storage interface {
	SaveURL(urlToSave string, alias string) (int64, error)
	GetURL(alias string) (string, error)
	// DeleteURL(alias string) error // Раскомментируйте если нужно
}

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExist    = errors.New("URL already exists")
)

//go:generate mockery --name=Storage --dir=. --output=./mocks --with-expecter
