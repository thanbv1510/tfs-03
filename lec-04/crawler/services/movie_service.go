package services

import (
	"bytes"
	"crawler/entities"
	"crawler/helpers"
	"crawler/models"
	"crawler/repositories"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var sugar = helpers.GetSugar()

func GetAllMovieLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" && n.Parent.Data == "td" && len(n.Attr) == 1 {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, fmt.Sprintf("https://imdb.com/%s", a.Val))
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = GetAllMovieLinks(links, c)
	}

	return links
}

func GetRawHTML(link string) string {
	resp, err := http.Get(link)
	if err != nil {
		sugar.Error(err.Error())
		return ""
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sugar.Error(err.Error())
		return ""
	}

	return string(respData)
}

func PushToChannel(linkChan chan string, links []string) {
	for i := 0; i < len(links); i++ {
		link := links[i]
		if len(strings.TrimSpace(link)) != 0 {
			sugar.Infof("==> [PUSH LINK] %s", link)
			linkChan <- link
		}
	}

	close(linkChan)
}

func ProcessData(linkChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	movies := make([]entities.Movie, 0)
	for {
		link, ok := <-linkChan
		if !ok {
			break
		}

		movie, err := GetMovieEntity(link)
		if err != nil {
			sugar.Error(err.Error())
			continue
		}

		sugar.Infof("==> [PROCESS MOVIE] %s", movie.String())
		movies = append(movies, movie)
		if len(movies) == 10 {
			repositories.InsertBatchMovie(movies)
			movies = make([]entities.Movie, 0)
		}
	}

	if len(movies) > 0 {
		repositories.InsertBatchMovie(movies)
	}
}

func GetMovieEntity(link string) (entities.Movie, error) {
	rawData := GetRawHTML(link)
	doc, err := html.Parse(strings.NewReader(rawData))
	if err != nil {
		sugar.Error(err.Error())
		return entities.Movie{}, err
	}
	allElementData := GetAllRawElementDetails([]string{}, doc)

	if len(allElementData) == 1 {
		movieDetail := GetData(allElementData[0])

		return movieDetail.MovieModel2MovieEntity(), nil
	}
	sugar.Errorf("DATA ERR with data: %s and link: %s", allElementData, link)
	return entities.Movie{}, errors.New("CANNOT GET DATA")
}

func GetAllRawElementDetails(data []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "script" {
		for _, a := range n.Attr {
			if a.Key == "type" && a.Val == "application/ld+json" {
				data = append(data, getRawNode(n))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		data = GetAllRawElementDetails(data, c)
	}

	return data
}

func getRawNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	err := html.Render(w, n)
	if err != nil {
		return ""
	}
	return buf.String()
}

func GetData(data string) (movieDetail models.MovieDetail) {
	tokenizer := html.NewTokenizer(strings.NewReader(data))
	for {
		tokenType := tokenizer.Next()
		if err := tokenizer.Err(); tokenType == html.ErrorToken && err == io.EOF {
			break
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if "script" == token.Data && strings.TrimSpace(token.Data) != "" {
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					err := json.Unmarshal([]byte(tokenizer.Token().Data), &movieDetail)
					if err != nil {
						sugar.Error(err.Error())
					}

					return
				}
			}
		}
	}
	return
}
