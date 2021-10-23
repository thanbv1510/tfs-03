package service

import (
	"awesomeProject/es"
	"awesomeProject/model"
	"awesomeProject/repository"
	"sync"
)

const url = "http://localhost:9200"

func SaveData(data <-chan model.Data, wg *sync.WaitGroup) {
	defer wg.Done()
	var result []model.Data
	client, _ := es.NewESClient(url)

	for {
		item, ok := <-data
		if !ok {
			break
		}
		result = append(result, item)
		if len(result) >= 1000 {
			repository.InsertBulk(result, client.Client)

			result = make([]model.Data, 0)
		}
	}

	if len(result) > 0 {
		repository.InsertBulk(result, client.Client)
	}
}

func FindByData(keyword string) []*model.Data {
	client, _ := es.NewESClient(url)
	return repository.FindData(keyword, client.Client)
}
