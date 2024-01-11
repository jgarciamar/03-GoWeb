package internal

import "errors"

var (
	ErrProductTitleAlreadyExists = errors.New("Product title already exists")
	ErrProductNotFound           = errors.New("Product not found")
)

// Product repository is an interface that represents
// a product repository

type ProductRepository interface {
	//Saves a proudct in the repository
	Save(product *Product) (err error)
	GetById(id int) (product Product, err error)
	GetAll() (products []Product, err error)
	Update(product *Product) (err error)
	Delete(id int) (err error)
}
