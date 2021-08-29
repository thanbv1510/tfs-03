package utils

import (
	"csv-import-es/models"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func CSVReader(path string, dataChan chan<- models.Review) {
	file, err := os.Open(path)
	if err != nil {
		// log err
		return
	}

	reader := csv.NewReader(file)
	reviewID := 1
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
		review.ID = strconv.Itoa(reviewID)
		reviewID++

		dataChan <- review
	}

	close(dataChan)
}
