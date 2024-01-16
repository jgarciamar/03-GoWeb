package service

import (
	"app/internal"
	"context"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {

	//Get stuff from the repository, use a empty context for now
	ctx := context.Background()
	products, err := s.rp.Get(ctx)

	if err != nil {
		return 0, err
	}

	totalTickets := len(products)

	return totalTickets, nil
}

func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(destination string) (tickets int, err error) {

	ticketCounter := 0
	ctx := context.Background()
	products, err := s.rp.Get(ctx)

	if err != nil {
		return 0, err
	}

	for _, v := range products {
		if v.Country == destination {
			ticketCounter++
		}
	}

	//Maybe check the validity of the destination?

	return ticketCounter, nil
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(destination string) (ticketPercentage float64, err error) {

	ticketCounter := 0
	ctx := context.Background()
	products, err := s.rp.Get(ctx)

	if err != nil {
		return 0, nil
	}

	totalAmount := len(products)

	for _, v := range products {
		if v.Country == destination {
			ticketCounter++
		}
	}

	ticketPercentage = float64(ticketCounter) / float64(totalAmount)

	return ticketPercentage, nil

}
