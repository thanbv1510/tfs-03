package main

import (
	"crawler/entities"
	"crawler/helpers"
	"crawler/services"
	"golang.org/x/net/html"
	"os"
	"strings"
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

	linkChan := make(chan string, 100)
	movieChan := make(chan entities.Movie, 100)
	doneChannel := make(chan bool)

	go services.PushToChannel(linkChan, movieLinks)
	go services.ProcessData(linkChan, movieChan)
	go services.SaveMovie(movieChan, doneChannel)

	_, _ = <-doneChannel
	sugar.Infof("==> DONE :) ")
}
