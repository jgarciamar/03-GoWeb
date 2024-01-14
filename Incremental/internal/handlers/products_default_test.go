package handlers_test

import (
	"clase-02/internal"
	"clase-02/internal/handlers"
	"clase-02/internal/repository"
	"clase-02/internal/service"
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
)

func TestProductsDefault_GetProducts(t *testing.T) {

	t.Run("Sucess 01 - Should return all tasks", func(t *testing.T) {

		// Arrange

		db := map[int]internal.Product{
			1: {Id: 1, Name: "Product 1", Quantity: 10, CodeValue: "code1", IsPublished: true, Expiration: "2021-12-31", Price: 100.0},
			2: {Id: 2, Name: "Product 2", Quantity: 20, CodeValue: "code2", IsPublished: true, Expiration: "2021-12-30", Price: 200.0},
			3: {Id: 3, Name: "Product 3", Quantity: 30, CodeValue: "code3", IsPublished: true, Expiration: "2021-12-29", Price: 300.0},
			4: {Id: 4, Name: "Product 4", Quantity: 40, CodeValue: "code4", IsPublished: true, Expiration: "2021-12-28", Price: 400.0},
		}

		rp := repository.NewProductMap(db, len(db))

		sv := service.NewProductDefault(rp)

		hd := handlers.NewDefaultProducts(sv)

		hdFunc := hd.GetAll()

		//Act

		//No parameter or query needed yet so we just pass it without any context.

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		res := httptest.NewRecorder()

		hdFunc(res, req)

		//Assert

		//Check if the status code is 200

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{
			"data": [
				{"id":1,"name":"Product 1","quantity":10,"code_value":"code1","is_published":true,"expiration":"2021-12-31","price":100},
				{"id":2,"name":"Product 2","quantity":20,"code_value":"code2","is_published":true,"expiration":"2021-12-30","price":200},
				{"id":3,"name":"Product 3","quantity":30,"code_value":"code3","is_published":true,"expiration":"2021-12-29","price":300},
				{"id":4,"name":"Product 4","quantity":40,"code_value":"code4","is_published":true,"expiration":"2021-12-28","price":400}
			],
			"message": "products fetched"
		}`
		require.Equal(t, expectedCode, res.Code, "Status code should be 200")
		require.Equal(t, expectedHeader, res.Header(), "Header should be the same")
		require.JSONEq(t, expectedBody, res.Body.String(), "Body should be the same")
	})
}

func TestProductsDefault_GetProductByID(t *testing.T) {

	db := map[int]internal.Product{
		1: {Id: 1, Name: "Product 1", Quantity: 10, CodeValue: "code1", IsPublished: true, Expiration: "2021-12-31", Price: 100.0},
		2: {Id: 2, Name: "Product 2", Quantity: 20, CodeValue: "code2", IsPublished: true, Expiration: "2021-12-30", Price: 200.0},
	}

	productID := 1

	rp := repository.NewProductMap(db, len(db))
	sv := service.NewProductDefault(rp)
	hd := handlers.NewDefaultProducts(sv)

	hdFunc := hd.GetById()

	//Manually add the urlParam to the request with the id of the product we want to get

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	res := httptest.NewRecorder()

	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams.Add("id", strconv.Itoa(productID))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

	hdFunc(res, req)

	expectedCode := http.StatusOK
	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
	expectedBody := `{
		"data": {
			"id":1,
			"name":"Product 1",
			"quantity":10,
			"code_value":"code1",
			"is_published":true,
			"expiration":"2021-12-31",
			"price":100
		},
		"message": "Product found"
	}`

	require.Equal(t, expectedCode, res.Code, "Status code should be 200")
	require.Equal(t, expectedHeader, res.Header(), "Header should be the same")
	require.JSONEq(t, expectedBody, res.Body.String(), "Body should be the same")

}

func TestProductsDefault_CreateProduct(t *testing.T) {

	//Arrange

	db := make(map[int]internal.Product, 0)

	rp := repository.NewProductMap(db, len(db))
	sv := service.NewProductDefault(rp)
	hd := handlers.NewDefaultProducts(sv)
	hdFunc := hd.Create()

	productJSONReader := strings.NewReader(`{
		"name":"Product 1",
		"quantity":10,
		"code_value":"code1",
		"is_published":true,
		"expiration":"2021-12-31",
		"price":100
	}`)

	req := httptest.NewRequest(http.MethodPost, "/products", productJSONReader)
	res := httptest.NewRecorder()

	// Act

	hdFunc(res, req)

	// Assert

	expectedCode := http.StatusCreated
	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
	expectedBody := `{
		"data": {
			"id":1,
			"name":"Product 1",
			"quantity":10,
			"code_value":"code1",
			"is_published":true,
			"expiration":"2021-12-31",
			"price":100
		},
		"message": "Product created"
	}`

	require.Equal(t, expectedCode, res.Code, "Status code should be 201")
	require.Equal(t, expectedHeader, res.Header(), "Header should be the same")
	require.JSONEq(t, expectedBody, res.Body.String(), "Body should be the same")
}

func TestProductsDefault_DeleteProduct(t *testing.T) {

	db := map[int]internal.Product{
		1: {Id: 1, Name: "Product 1", Quantity: 10, CodeValue: "code1", IsPublished: true, Expiration: "2021-12-31", Price: 100.0},
	}
	rp := repository.NewProductMap(db, len(db))
	sv := service.NewProductDefault(rp)
	hd := handlers.NewDefaultProducts(sv)
	hdFunc := hd.Delete()

	//Assert

	expectedCode := http.StatusOK
	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
	expectedBody := `{
		"message": "Product deleted"
	}`

	//Act

	productID := 1

	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	res := httptest.NewRecorder()

	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams.Add("id", strconv.Itoa(productID))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

	hdFunc(res, req)

	require.Equal(t, expectedCode, res.Code, "Status code should be 200")
	require.Equal(t, expectedHeader, res.Header(), "Header should be the same")
	require.JSONEq(t, expectedBody, res.Body.String(), "Body should be the same")

	//Also check the db maybe?

	require.Len(t, db, 0, "Database should be empty after the delete")
}
