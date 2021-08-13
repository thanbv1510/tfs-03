package repositories

import (
	"crawler/entities"
	"crawler/helpers"
	"sync"
)

func InsertBatchMovie(movies []entities.Movie) {
	var mu sync.Mutex
	mu.Lock()
	helpers.DBConn().CreateInBatches(movies, len(movies))
	mu.Unlock()
}
