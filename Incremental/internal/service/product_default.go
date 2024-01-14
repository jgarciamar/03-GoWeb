package service

import (
	"clase-02/internal"
	"fmt"
)

func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

type ProductDefault struct {
	rp internal.ProductRepository
}

func (p *ProductDefault) Save(product *internal.Product) (err error) {

	if (*product).Name == "" {
		return fmt.Errorf("%w: name", internal.ErrFieldRequired)
	}

	if product.Quantity <= 0 {
		return fmt.Errorf("%w: quantity", internal.ErrFieldRequired)
	}

	if product.CodeValue == "" {
		return fmt.Errorf("%w: code_value", internal.ErrFieldRequired)
	}

	if product.Expiration == "" {
		return fmt.Errorf("%w: expiration", internal.ErrFieldRequired)
	}

	if product.Price <= 0 {
		return fmt.Errorf("%w: price", internal.ErrFieldRequired)
	}

	err = p.rp.Create(product)

	// We now check errors returned by the repo

	if err != nil {
		switch err {
		case internal.ErrProductTitleAlreadyExists:
			return fmt.Errorf("%w: title", internal.ErrProductAlreadyExists)
		}
		return
	}

	return nil
}

func (p *ProductDefault) GetAll() ([]internal.Product, error) {

	products, err := p.rp.GetAll()

	if err != nil {
		fmt.Println(err)
		return products, err
	}

	return products, nil

}

func (p *ProductDefault) GetById(id int) (product internal.Product, err error) {

	product, err = p.rp.GetById(id)

	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		default:
			fmt.Println(err)
			err = fmt.Errorf("Error in Service")
		}

		return
	}
	return
}

func ValidateProduct(product *internal.Product) (err error) {
	return nil
}

func (p *ProductDefault) Update(product *internal.Product) (err error) {
	if err = ValidateProduct(product); err != nil {
		return
	}

	err = p.rp.Update(product)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}
	return
}

func (p *ProductDefault) Delete(id int) (err error) {
	err = p.rp.Delete(id)

	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}
	return

}
