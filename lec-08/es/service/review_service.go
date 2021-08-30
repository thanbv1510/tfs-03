package service

import (
	"csv-import-es/es"
	"csv-import-es/models"
	"csv-import-es/repository"
	"sync"
)

const url = "http://localhost:9200"

func SaveReview(data <-chan models.Review, wg *sync.WaitGroup) {
	defer wg.Done()
	var reviews []models.Review
	client, _ := es.NewESClient(url)

	for {
		review, ok := <-data
		if !ok {
			break
		}
		reviews = append(reviews, review)
		if len(reviews) >= 10000 {
			repository.InsertBulkReview(reviews, client.Client)

			reviews = make([]models.Review, 0)
		}
	}

	if len(reviews) > 0 {
		repository.InsertBulkReview(reviews, client.Client)
	}
}

func FindByBodyReview(keyword string) []*models.Review {
	client, _ := es.NewESClient(url)
	return repository.FindByBodyReview(keyword, client.Client)
}
