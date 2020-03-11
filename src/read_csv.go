package src

import (
	"encoding/csv"
	"io"
	"log"
)

const (
	buffer = 1000
	freq   = 100
)

func ReadData(reader *csv.Reader) ([]InputRecord, error) {
	records := make([]InputRecord, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		inrecord, err1 := FromCSV(record)

		if err1 != nil {
			continue
		}

		records = append(records, inrecord)
	}
	return records, nil
}


func WriteOutputData(file io.Writer, records []OutputRecord) {
	w := csv.NewWriter(file)
	var err error
	var i int
	for _, record := range records {
		if err = w.Write(record.CSV()); err != nil {
			return
		}

		i++
		if i%freq == 0 {
			w.Flush()
			if err = w.Error(); err != nil {
				return
			}
			i = 0
		}
	}
	if i > 0 {
		w.Flush()
		if err = w.Error(); err != nil {
			return
		}
	}
}