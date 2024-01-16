package service_test

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("Sucess 01:  Get total tickets", func(t *testing.T) {
		// arrange

		db := map[int]internal.TicketAttributes{
			1: {Name: "John", Email: "juan@gmail.com", Country: "USA", Hour: "10:00", Price: 100},
			2: {Name: "Camille", Email: "jorge@gmail.com", Country: "Colombia", Hour: "10:00", Price: 200},
			3: {Name: "Jorge", Email: "", Country: "Colombia", Hour: "10:00", Price: 800},
		}

		rp := repository.NewRepositoryTicketMap(db, len(db))

		sv := service.NewServiceTicketDefault(rp)

		expectedTotal := 3

		// act
		total, err := sv.GetTotalAmountTickets()

		assert.NoError(t, err, "There should be no error")
		assert.Equal(t, expectedTotal, total, "The total should be 3")

	})
}

func TestServiceTicketDefault_GetTicketsAmountByDestinationCountry(t *testing.T) {
	t.Run("Sucess 01: To get percentage tickets by destination country", func(t *testing.T) {
		// arrange

		db := map[int]internal.TicketAttributes{
			1: {Name: "John", Email: "juan@gmail.com", Country: "USA", Hour: "10:00", Price: 100},
			2: {Name: "Camille", Email: "jorge@gmail.com", Country: "Colombia", Hour: "10:00", Price: 200},
			3: {Name: "Jorge", Email: "jorge@gmail.com", Country: "Colombia", Hour: "10:00", Price: 800},
			4: {Name: "Jose", Email: "jorge@gmail.com", Country: "Colombia", Hour: "10:00", Price: 800},
		}

		rp := repository.NewRepositoryTicketMap(db, len(db))

		sv := service.NewServiceTicketDefault(rp)

		// act

		expectedTicketAmount := 3
		ticketAmount, err := sv.GetTicketsAmountByDestinationCountry("Colombia")

		assert.NoError(t, err, "There should be no error")
		assert.Equal(t, expectedTicketAmount, ticketAmount, "The ticket amount should be 3")

	})
}

func TestServiceTicketDefault_GetPercentageTicketsByDestinationCountry(t *testing.T) {
	t.Run("Sucess 03: To get percentage tickets by destination country", func(t *testing.T) {
		// arrange

		db := map[int]internal.TicketAttributes{
			1: {Name: "John", Email: "juan@gmail.com", Country: "USA", Hour: "10:00", Price: 100},
			2: {Name: "Camille", Email: "jorge@gmail.com", Country: "Colombia", Hour: "10:00", Price: 200},
			3: {Name: "Jorge", Email: "jorge@gmail.com", Country: "Colombia", Hour: "10:00", Price: 800},
			4: {Name: "Jose", Email: "jorge@gmail.com", Country: "Colombia", Hour: "10:00", Price: 800},
		}

		rp := repository.NewRepositoryTicketMap(db, len(db))

		sv := service.NewServiceTicketDefault(rp)

		// act

		expectedPercentage := float64(3) / float64(4)
		ticketPercentage, err := sv.GetPercentageTicketsByDestinationCountry("Colombia")

		assert.NoError(t, err, "There should be no error")
		assert.Equal(t, expectedPercentage, ticketPercentage, "Percentage should be 0.75")

	})

}
