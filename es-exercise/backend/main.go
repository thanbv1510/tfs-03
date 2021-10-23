package main

import (
	"awesomeProject/model"
	"awesomeProject/server"
	"awesomeProject/service"
	"awesomeProject/utils"
	"fmt"
	"sync"
)

func main() {

	path := "./lme-corpus.csv"
	dataChan := make(chan model.Data)
	go func() {
		utils.CSVReader(path, dataChan)
	}()

	numGoroutine := 10
	var wg sync.WaitGroup
	wg.Add(numGoroutine)
	for i := 0; i < numGoroutine; i++ {
		service.SaveData(dataChan, &wg)
	}
	wg.Wait()
	fmt.Println("Done!")

	server.InitServer()
}
