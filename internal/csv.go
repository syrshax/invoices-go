package internal

import (
	"encoding/csv"
	"os"
)

type CSVRow struct {
	Name       string
	NationalID string
	Address    string
}

func ReadCSV(filepath string) ([]CSVRow, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	var rows []CSVRow

	idx := 1
	if len(records) <= 1 {
		return rows, nil
	}

	for i := idx; i < len(records); i++ {
		r := records[i]
		row := CSVRow{}
		if len(r) > 0 {
			row.Name = r[0]
		}
		if len(r) > 1 {
			row.NationalID = r[1]
		}
		if len(r) > 2 {
			row.Address = r[2]
		}
		rows = append(rows, row)
	}
	return rows, nil
}
