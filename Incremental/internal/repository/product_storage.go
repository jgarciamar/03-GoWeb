package repository

import (
	"clase-02/internal"
	"clase-02/internal/storage"
)

type RepositoryProductStore struct {
	st     storage.StorageProduct
	lastID int
}

func NewRepositoryProductStore(st storage.StorageProduct, lastID int) *RepositoryProductStore {
	return &RepositoryProductStore{
		st:     st,
		lastID: lastID,
	}
}

func (r *RepositoryProductStore) GetById(id int) (product internal.Product, err error) {

	ps, err := r.st.ReadAll()

	if err != nil {
		return
	}

	foundProduct := false

	for _, v := range ps {
		if v.Id == id {
			product = v
			foundProduct = true
			return
		}
	}

	if !foundProduct {
		err = internal.ErrProductDoesNotExist
	}
	return
}

func (r *RepositoryProductStore) Create(product *internal.Product) (err error) {

	ps, err := r.st.ReadAll()

	if err != nil {
		return err
	}

	for _, v := range ps {

		if v.Name == product.Name {
			return internal.ErrProductTitleAlreadyExists
		}

	}

	(*r).lastID++

	(*product).Id = (*r).lastID

	ps = append(ps, *product)

	err = r.st.WriteAll(ps)

	if err != nil {
		return err
	}

	return nil

}

func (r *RepositoryProductStore) GetAll() (products []internal.Product, err error) {

	ps, err := r.st.ReadAll()

	if err != nil {
		return products, err
	}

	return ps, nil

}

func (r *RepositoryProductStore) Update(product *internal.Product) (err error) {

	ps, err := r.st.ReadAll()

	if err != nil {
		return err
	}

	var found bool
	var pIndex int
	for i, v := range ps {
		if v.Id == product.Id {
			found = true
			pIndex = i
			break
		}
	}

	if !found {
		return internal.ErrProductNotFound
	}

	// We replace the product, this seems very unsafe and may be
	// needed to recheck how to this better than now
	ps[pIndex] = *product

	if err := r.st.WriteAll(ps); err != nil {
		return err
	}

	return nil

}

func (r *RepositoryProductStore) Delete(id int) (err error) {
	ps, err := r.st.ReadAll()

	if err != nil {
		return err
	}

	//Search the id to see if exists
	var found bool
	index := -1

	for i, v := range ps {
		if v.Id == id {
			found = true
			index = i
			break
		}
	}

	if !found {
		return internal.ErrProductNotFound
	}

	// We delete the item from the slice

	ps = append(ps[:index], ps[index+1:]...)

	err = r.st.WriteAll(ps)

	if err != nil {
		return err
	}

	return nil

}
