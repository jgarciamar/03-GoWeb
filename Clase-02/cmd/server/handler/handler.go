package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type MyHandler struct {
	data map[string]string
}

func NewHandler() *MyHandler {
	return &MyHandler{
		data: map[string]string{
			"1": "Garcia",
			"2": "Martinez",
			"3": "Jose",
		},
	}
}

func (h *MyHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	}
}

type MyResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (h *MyHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")

		name, ok := h.data[id]

		if !ok {
			code := http.StatusNotFound
			body := MyResponse{Message: "User not found", Data: nil}

			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(body)
		}

		w.Write([]byte(name))
	}
}
