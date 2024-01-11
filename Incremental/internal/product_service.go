package internal

import "errors"

var (
	ErrFieldRequired        = errors.New("Field Required!!!")
	ErrFieldQuality         = errors.New("Field quality")
	ErrProductAlreadyExists = errors.New("Product already exists")
)

type ProductService interface {
	Save(product *Product) (err error)
	GetById(id int) (product Product, err error)
	GetAll() (products []Product, err error)
	Update(product *Product) (err error)
	Delete(id int) (err error)
}
