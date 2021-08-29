package main

import (
	"csv-import-es/models"
	"csv-import-es/service"
	"csv-import-es/utils"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start date and time is: ", time.Now())

	path := "../train.csv"
	dataChan := make(chan models.Review)
	go func() {
		utils.CSVReader(path, dataChan)
	}()

	numGoroutineReview := 10000
	var wg sync.WaitGroup
	wg.Add(numGoroutineReview)
	for i := 0; i < numGoroutineReview; i++ {
		service.SaveReview(dataChan, &wg)
	}
	wg.Wait()

	fmt.Println("End date and time is: ", time.Now())
}
