package utils

import (
	"csv-import-mysql/models"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func CSVReader(path string, dataChan chan<- models.Review) {
	file, err := os.Open(path)
	if err != nil {
		// log err
		return
	}

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		var review models.Review
		for i, v := range record {
			switch i {
			case 0:
				review.Type = v
			case 1:
				review.Title = v
			case 2:
				review.Body = v
			}
		}
		dataChan <- review
	}

	close(dataChan)
}
