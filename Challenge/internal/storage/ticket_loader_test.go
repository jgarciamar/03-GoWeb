package storage_test

import (
	"app/internal"
	"app/internal/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromCSV(t *testing.T) {

	loader := storage.NewLoaderTicketCSV("../../tickets.csv")

	mapDB, err := loader.Load()

	if err != nil {
		t.Error(err)
	}

	if len(mapDB) == 0 {
		t.Error("The map is empty")
	}

	assert.Equal(
		t,
		mapDB[1],
		internal.TicketAttributes{
			Name:    "Tait Mc Caughan",
			Email:   "tmc0@scribd.com",
			Country: "Finland",
			Hour:    "17:11",
			Price:   785},
		"LoadFromCV did not map correctly",
	)
}
