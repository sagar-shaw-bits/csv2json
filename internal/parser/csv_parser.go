package parser

import (
	"encoding/csv"
	"io"
)

func ParseCSV(r io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(r)

	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var records []map[string]string

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		record := make(map[string]string)
		for i, value := range row {
			record[headers[i]] = value
		}
		records = append(records, record)
	}

	return records, nil
}
