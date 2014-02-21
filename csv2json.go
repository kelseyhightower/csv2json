package csv2json

import (
	"encoding/csv"
	"encoding/json"
)

type famousGopher struct {
	Name  string
	Date  string
	Title string
}

func Convert(r io.Reader) ([]byte, error) {
	gophers := make([]famousGopher, 0)
	csvReader := csv.NewReader(r)
	csvReader.TrimLeadingSpace = true
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		gopher := famousGopher{
			Name:  record[0],
			Date:  record[1],
			Title: record[2],
		}
		gophers = append(gophers, gopher)
	}
	data, err := json.MarshalIndent(&gophers, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
