package service

import (
	"csv-import-mysql/database"
	"csv-import-mysql/models"
	"csv-import-mysql/repository"
	"fmt"
	"sync"
)

func SaveReview(data <-chan models.Review, wg *sync.WaitGroup) {
	defer wg.Done()
	var reviews []models.Review
	for {
		review, ok := <-data
		if !ok {
			break
		}
		reviews = append(reviews, review)
		if len(reviews) >= 10000 {
			sqlDB, _ := database.DBConn()
			err := repository.InsertBatchReview(reviews, sqlDB)
			if err != nil {
				fmt.Println(err)
			}

			reviews = make([]models.Review, 0)
		}
	}

	if len(reviews) > 0 {
		sqlDB, _ := database.DBConn()
		err := repository.InsertBatchReview(reviews, sqlDB)
		if err != nil {
			fmt.Println(err)
		}
	}
}
