package handler

import (
	"app/internal"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TicketDefault struct {
	sv internal.ServiceTicket
}

func NewTicketDefault(sv internal.ServiceTicket) *TicketDefault {
	return &TicketDefault{
		sv: sv,
	}
}

func (t *TicketDefault) GetTicketAmountByCountry() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		destination := chi.URLParam(r, "dest")

		if destination == "" {
			http.Error(w, "Missing destination", http.StatusBadRequest)
			return
		}

		ticketAmount, err := t.sv.GetTicketsAmountByDestinationCountry(destination)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"amount": ticketAmount,
		})
	}
}

func (t *TicketDefault) GetAverageTicketByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		destination := chi.URLParam(r, "dest")

		if destination == "" {
			http.Error(w, "Missing destination", http.StatusBadRequest)
			return
		}

		ticketPercentage, err := t.sv.GetPercentageTicketsByDestinationCountry(destination)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"country":    destination,
			"percentage": ticketPercentage,
		})
	}
}
