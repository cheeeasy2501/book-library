package book

import "github.com/google/uuid"

type BookInterface interface {
	GetByUUID(id uuid.UUID) (*Book, error)
	GetAll() ([]*Book, error)
	Create(*Book) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(id uuid.UUID) error
}

type BookRepository struct {
}

func (br *BookRepository) GetByUUID(id uuid.UUID) (*Book, error) {
	var book *Book

	return book, nil
}

func (br *BookRepository) GetAll() ([]*Book, error) {
	var books []*Book

	return books, nil
}

func (br *BookRepository) Create(book *Book) (*Book, error) {

	return book, nil
}

func (br *BookRepository) Update(book *Book) (*Book, error) {

	return book, nil
}

func (br *BookRepository) Delete(id uuid.UUID) error {

	return nil
}
