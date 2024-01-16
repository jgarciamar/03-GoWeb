package repository_test

import (
	"app/internal/repository"
	"app/internal/storage"
	"context"
	"testing"
)

func TestTicketMap(t *testing.T) {

	t.Run("Sucess 01: Get all tickets", func(t *testing.T) {

		storage := storage.NewLoaderTicketCSV("../../tickets.csv")
		db, err := storage.Load()

		if err != nil {
			t.Error(err)
		}

		rp := repository.NewRepositoryTicketMap(db, len(db))
		ctx := context.Background()

		testProduct, err := rp.Get(ctx)
		if err != nil {
			t.Error(err)
		}

		if len(testProduct) == 0 {
			t.Error("The map is empty")
		}
	})
}
