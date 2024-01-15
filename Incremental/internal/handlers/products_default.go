package handlers

import (
	"clase-02/internal"
	"clase-02/web"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"

	"github.com/go-chi/chi"
)

func NewDefaultProducts(sv internal.ProductService) *DefaultProducts {
	return &DefaultProducts{
		sv: sv,
	}
}

type DefaultProducts struct {
	sv internal.ProductService
}

type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

//GetBtId returns a product by id

func (d *DefaultProducts) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			web.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

		product, err := d.sv.GetById(id)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				web.Error(w, http.StatusNotFound, "Product not found")
			default:
				//fmt.Println(err)
				web.Error(w, http.StatusInternalServerError, "Internal error in handling the request")
			}
			return
		}

		data := ProductJSON{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		//add an err check

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Product found",
			"data":    data,
		})
	}
}

func (d *DefaultProducts) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body BodyRequestProductJSON

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			web.Error(w, http.StatusBadRequest, "Invalid body")
			return
		}

		// Serialize internal.Product

		product := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		//Checking errors on the service

		if err := d.sv.Save(&product); err != nil {
			web.Error(w, http.StatusBadRequest, "unable to create the product")
			return
		}

		//Response

		data := ProductJSON{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Product created",
			"data":    data,
		})
	}
}
func (d *DefaultProducts) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Fetch all products from the service
		products, err := d.sv.GetAll()
		if err != nil {
			web.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// Serialize products to JSON
		var responseData []ProductJSON
		for _, product := range products {
			data := ProductJSON{
				Id:          product.Id,
				Name:        product.Name,
				Quantity:    product.Quantity,
				CodeValue:   product.CodeValue,
				IsPublished: product.IsPublished,
				Expiration:  product.Expiration,
				Price:       product.Price,
			}
			responseData = append(responseData, data)
		}

		// Response
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "products fetched",
			"data":    responseData,
		})
	}
}

func (d *DefaultProducts) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			web.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

		bytes, err := io.ReadAll(r.Body)

		if err != nil {
			web.Error(w, http.StatusBadRequest, "Invalid body")
			return
		}

		var bodyMap map[string]any

		if err := json.Unmarshal(bytes, &bodyMap); err != nil {

			web.Error(w, http.StatusBadRequest, "Invalid body")
			return
		}

		PossibleRequestFields := []string{"name"}

		if err := ValidateKeyExistence(bodyMap, PossibleRequestFields); err != nil {

			web.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		var body BodyRequestProductJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			web.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		product := internal.Product{
			Id:          id,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		if err := d.sv.Update(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				web.Error(w, http.StatusNotFound, "Product not found")
			case errors.Is(err, internal.ErrFieldRequired):
				web.Error(w, http.StatusBadRequest, "Invalid body")
			default:
				web.Error(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		data := ProductJSON{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Product updated",
			"data":    data,
		})
	}

}

// Delete delates a product from the db
// It returns a 204 if the product was deleted successfully
func (d *DefaultProducts) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			web.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

		if err := d.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				web.Error(w, http.StatusNotFound, "product not found")
			default:
				web.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "Product deleted",
		})
	}
}

// UpdatePartial updates a movie
func (d *DefaultProducts) UpdatePartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			web.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		product, err := d.sv.GetById(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				web.Error(w, http.StatusNotFound, "product not found")
			default:
				web.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// process
		// - serialize Product into request
		reqBody := BodyRequestProductJSON{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		// - get body
		if err := request.JSON(r, &reqBody); err != nil {
			web.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - serialize
		product = internal.Product{
			Id:          id,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}

		// - update
		if err := d.sv.Update(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				web.Error(w, http.StatusNotFound, "movie not found")
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				web.Error(w, http.StatusBadRequest, "invalid body")
			default:
				web.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - deserialize MovieJSON
		data := ProductJSON{
			Id:          id,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product updated",
			"data":    data,
		})

	}
}
func ValidateKeyExistence(mp map[string]any, keys []string) (err error) {
	for _, k := range keys {
		if _, ok := mp[k]; !ok {
			return fmt.Errorf("Key %s not found", k)
		}
	}
	return
}
