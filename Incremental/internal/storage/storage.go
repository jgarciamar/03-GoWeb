package storage

import "clase-02/internal"

type StorageProduct interface {
	ReadAll() (products []internal.Product, err error)
	WriteAll(products []internal.Product) (err error)
}
