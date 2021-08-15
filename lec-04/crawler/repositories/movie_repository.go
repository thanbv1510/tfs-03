package repositories

import (
	"crawler/entities"
	"crawler/helpers"
	"sync"
)

var sugar = helpers.GetSugar()

func InsertBatchMovie(movies []entities.Movie) {
	var mu sync.Mutex
	mu.Lock()
	helpers.DBConn().CreateInBatches(movies, len(movies))
	sugar.Infof("==> [SAVE MOVIE] inserted %d element", len(movies))
	mu.Unlock()
}
