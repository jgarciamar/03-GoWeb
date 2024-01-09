package main

import (
	"net/http"
	"routing/cmd/server/handler"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	h := handler.NewHandler()

	r.Get("/", h.Get())

	r.Get("/users/{userId}", h.GetById())
	http.ListenAndServe(":8080", r)
}
