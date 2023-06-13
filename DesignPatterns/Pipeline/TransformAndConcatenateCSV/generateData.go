package main

import (
  "encoding/csv"
  "io"
  "os"
  "errors"
)


func generateData(filename string) (<-chan []string, error) {
	f, err := os.Open(filename)
	if err!= nil {
        return nil, errors.New(filename)
    }

	output := make(chan []string)
	go func() {
		cr := csv.NewReader(f)
		cr.FieldsPerRecord = 3

		for {
			record, err := cr.Read()
            if err == io.EOF {
                close(output)
                return
            } 

			output <- record
			}
		}()
		return output, nil
}