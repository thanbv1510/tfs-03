package main

import (
	"crawler/entities"
	"crawler/models"
	"crawler/utils"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	err := utils.DBConn().AutoMigrate(&entities.Movie{}, &entities.AggregateRating{}, &entities.Actor{}, &entities.Creator{}, &entities.Director{}, &entities.Genre{})
	if err != nil {
		fmt.Println("Cannot create all table!")
		panic(err)
	}

	link := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	data := utils.GetRawHTML(link)
	/*data, err := ioutil.ReadFile("imdb.html")*/

	doc, err := html.Parse(strings.NewReader(data))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	allElementData := utils.GetAllRawElementData([]string{}, doc)

	rawDataChan := make(chan models.MovieInfo, 100)
	doneChannel := make(chan bool)

	go utils.PushToChannel(rawDataChan, allElementData)
	for i := 0; i < 100; i++ {
		go utils.SaveToDB(rawDataChan, doneChannel)
	}

	_, _ = <-doneChannel

	/*data, _ := ioutil.ReadFile("batman.html")
	doc, err := html.Parse(strings.NewReader(string(data)))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	allElementData := utils.GetAllRawElementDetails([]string{}, doc)

	movieDetail := utils.GetData(allElementData[0])
	fmt.Println(movieDetail)
	aggregateRating := entities.AggregateRating{Type: movieDetail.AggregateRating.Type}
	var directorEntity []*entities.Director
	for _, v := range movieDetail.Directors {
		directorEntity = append(directorEntity, &entities.Director{Name: v.Name})
	}

	moVieEntity := entities.Movie{Name: movieDetail.Name, AggregateRating: aggregateRating, Director: directorEntity}

	if err != nil {
		fmt.Println(err)
		return
	}

	result := utils.DBConn().Create(&moVieEntity)
	fmt.Println(result)*/
}
