package main

import (
	"crawler/entities"
	"crawler/helpers"
	"crawler/services"
	"golang.org/x/net/html"
	"os"
	"strings"
	"sync"
)

func main() {
	sugar := helpers.GetSugar()

	// Auto create table
	err := helpers.DBConn().AutoMigrate(&entities.Movie{}, &entities.AggregateRating{}, &entities.Actor{}, &entities.Creator{}, &entities.Director{})
	if err != nil {
		sugar.Error("==> Cannot create table!")
		panic(err)
	}

	link := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	data := services.GetRawHTML(link)

	doc, err := html.Parse(strings.NewReader(data))
	if err != nil {
		sugar.Error(err.Error())
		os.Exit(1)
	}

	movieLinks := services.GetAllMovieLinks([]string{}, doc)
	sugar.Infof("==> Find %d link movies", len(movieLinks))
	if len(movieLinks) == 0 {
		sugar.Info("==> EXIT :( ")
		return
	}

	linkChan := make(chan string)
	go services.PushToChannel(linkChan, movieLinks)

	numGoroutine := 10
	var wg sync.WaitGroup
	wg.Add(numGoroutine)
	for i := 0; i < numGoroutine; i++ {
		go services.ProcessData(linkChan, &wg)
	}
	wg.Wait()

	sugar.Infof("==> DONE :) ")
}
