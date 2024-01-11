package repository

import "clase-02/internal"

type ProductMap struct {
	db     map[int]internal.Product
	lastId int
}

func NewProductMap(db map[int]internal.Product, lastId int) *ProductMap {
	//default config and values
	//...

	return &ProductMap{
		db:     db,
		lastId: lastId,
	}
}

func (p *ProductMap) Save(product *internal.Product) (err error) {

	for _, v := range (*p).db {

		if v.Name == product.Name {
			return internal.ErrProductTitleAlreadyExists
		}

	}

	(*p).lastId++

	(*product).Id = (*p).lastId

	// Store

	(*p).db[(*product).Id] = *product

	return nil
}

//GetAll() (products []*Product, err error)

func (p *ProductMap) GetAll() (products []internal.Product, err error) {

	var Products []internal.Product

	for _, p := range p.db {
		Products = append(Products, p)
	}

	return Products, nil

}

func (p *ProductMap) GetById(id int) (product internal.Product, err error) {
	product, ok := p.db[id]

	if !ok {
		err = internal.ErrProductNotFound
		return product, err
	}
	return
}

func (p *ProductMap) Update(product *internal.Product) (err error) {
	_, ok := p.db[product.Id]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}

	p.db[product.Id] = *product
	return
}

func (p *ProductMap) Delete(id int) (err error) {
	_, ok := p.db[id]

	if !ok {
		err = internal.ErrProductNotFound
		return
	}
	delete(p.db, id)
	return
}
