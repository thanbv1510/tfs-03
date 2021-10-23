package utils

import (
	"awesomeProject/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func CSVReader(path string, dataChan chan<- model.Data) {
	file, err := os.Open(path)
	if err != nil {
		// log err
		panic(err)
	}

	reader := csv.NewReader(file)
	dataID := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		var data model.Data
		for i, v := range record {
			switch i {
			case 0:
				data.Text = v
			case 1:
				data.FileName = v
			case 2:
				data.MEDShortTitle = v
			case 3:
				data.SourceCorpus = v
			case 4:
				data.Edn = v
			case 5:
				data.MS = v
			case 6:
				data.MEDData = v
			case 7:
				data.Area = v
			case 8:
				data.VP = v
			case 9:
				data.Genre = v
			case 10:
				data.Words = v
			}
		}
		data.ID = strconv.Itoa(dataID)
		dataID++

		dataChan <- data
	}

	close(dataChan)
}

func ResponseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(object); err != nil {
		return
	}
}
