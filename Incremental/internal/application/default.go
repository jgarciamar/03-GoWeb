package application

import (
	"clase-02/internal"
	"clase-02/internal/handler"
	"clase-02/internal/repository"
	"clase-02/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func NewDefaultHTTP(addr string) *DefaultHTTP {
	return &DefaultHTTP{
		addr: addr,
	}
}

type DefaultHTTP struct {
	addr string
}

func GetProductsFromJSON(jsonPath string) (map[int]internal.Product, error) {
	defaultFilePath := jsonPath

	if defaultFilePath == "" {
		defaultFilePath = "products.json"
	}

	content, err := os.ReadFile(defaultFilePath)
	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return map[int]internal.Product{}, err
	}

	var JSONProducts []internal.Product

	err = json.Unmarshal(content, &JSONProducts)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return map[int]internal.Product{}, err
	}

	productMap := make(map[int]internal.Product)

	for _, product := range JSONProducts {
		productMap[product.Id] = product
	}

	fmt.Println("Products from JSONFile have been loaded")

	return productMap, nil
}

func (h *DefaultHTTP) Run() (err error) {

	JSONProductsMap, err := GetProductsFromJSON("products.json")

	if err != nil {
		fmt.Println(err)
		return err
	}

	rp := repository.NewProductMap(JSONProductsMap, len(JSONProductsMap)-1)

	sv := service.NewProductDefault(rp)

	hd := handler.NewDefaultProducts(sv)

	rt := chi.NewRouter()

	rt.Route("/products", func(rt chi.Router) {

		rt.Post("/", hd.Create())
		rt.Get("/", hd.GetAll())
		rt.Get("/{id}", hd.GetById())
		rt.Put("/{id}", hd.Update())
		rt.Delete("/{id}", hd.Delete())
		rt.Patch("/{id}", hd.UpdatePartial())
	})

	err = http.ListenAndServe(h.addr, rt)
	return

}
