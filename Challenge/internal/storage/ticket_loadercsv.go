package storage

import (
	"app/internal"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {

	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (l *LoaderTicketCSV) Load() (t map[int]internal.TicketAttributes, err error) {
	// open the file
	f, err := os.Open(l.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records

	t = make(map[int]internal.TicketAttributes)

	for {
		record, err := r.Read()

		if err != nil {

			if err != io.EOF {
				return nil, err
			} else {
				break
			}
		}

		// serialize the record

		id, err := strconv.Atoi(record[0])

		if err != nil {
			return nil, fmt.Errorf("error parsing")
		}

		price, err := strconv.ParseFloat(record[5], 64)

		if err != nil {
			return nil, fmt.Errorf("error parsing price: %v", err)
		}

		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}
		t[id] = ticket
	}

	//fmt.Println(t)

	return t, nil
}
