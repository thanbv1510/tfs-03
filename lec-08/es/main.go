package main

import (
	"csv-import-es/service"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	/*
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
	*/
	keyword := "a JVC"
	reviews := service.FindByBodyReview(keyword)
	fmt.Printf("Searching with the keyword '%s' of body got %d results\n", keyword, len(reviews))
	fmt.Printf("Find took %f (seconds)\n", time.Since(start).Seconds())
}
