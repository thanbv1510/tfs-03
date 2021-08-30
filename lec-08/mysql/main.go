package main

import (
	"csv-import-mysql/service"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	/*
		dbSql, _ := database.DBConn()
		_ = repository.CreateReviewTable(dbSql)

		path := "../train.csv"
		dataChan := make(chan models.Review)
		go func() {
			utils.CSVReader(path, dataChan)
		}()

		numGoroutineReviews := 1000
		var wg sync.WaitGroup
		wg.Add(numGoroutineReviews)
		for i := 0; i < numGoroutineReviews; i++ {
			service.SaveReview(dataChan, &wg)
		}
		wg.Wait()
	*/

	keyword := "a JVC"
	reviews := service.FindByBodyReview(keyword)
	fmt.Printf("Searching with the keyword '%s' of body got %d results\n", keyword, len(reviews))
	fmt.Printf("Find took %f (seconds)\n", time.Since(start).Seconds())
}
