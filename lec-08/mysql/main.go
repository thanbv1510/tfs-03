package main

import (
	"csv-import-mysql/database"
	"csv-import-mysql/models"
	"csv-import-mysql/repository"
	"csv-import-mysql/service"
	"csv-import-mysql/utils"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start date and time is: ", time.Now())

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
	fmt.Println("End date and time is: ", time.Now())
	fmt.Println("Done!")
}
