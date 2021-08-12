package utils

import (
	"bytes"
	"crawler/entities"
	"crawler/models"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAllRawElementData(data []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "tr" {
		data = append(data, getRawNode(n))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		data = GetAllRawElementData(data, c)
	}

	return data
}

func GetRawHTML(link string) string {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(respData)
}

/*func GetRawData(n *html.Node, dataChan chan *html.Node) {
	for {
		if n.Type == html.ElementNode && n.Data == "tr" {
			dataChan <- n
		}

		if c := n.FirstChild; c != nil {
			n = c.NextSibling
		}

		//for c := n.FirstChild; c != nil; c = c.NextSibling {
		//	GetRawData(c, dataChan)
		//}
	}
}*/

/*func ProcessRawData(dataChan chan *html.Node) {
	counter := 1
	note, isContinue := <-dataChan
	if isContinue {
		GetTextData(note)
		counter++
	} else {
		fmt.Println("read all data")
		return
	}
}*/

func getRawNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	err := html.Render(w, n)
	if err != nil {
		return ""
	}
	return buf.String()
}

/*func GetTextData(n *html.Node) {
	tokenizer := html.NewTokenizer(strings.NewReader(getRawNode(n)))

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if "a" == token.Data || "strong" == token.Data && strings.TrimSpace(token.Data) != "" {
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					fmt.Println(tokenizer.Token().Data)
					//for _, arr := range token.Attr {
					//	if arr.Key == "href" {
					//		fmt.Println(arr.Val)
					//	}
					//}
				}
			}
		}

		if token := tokenizer.Token(); tokenType == html.EndTagToken && "tr" == token.Data {
			break
		}

	}
}*/

func PushToChannel(ch chan models.MovieInfo, rawDataChan []string) {
	for i := 0; i < len(rawDataChan); i++ {
		tokenizer := html.NewTokenizer(strings.NewReader(rawDataChan[i]))
		var Name string
		var UrlMovie string
		var UrlImage string
		var Rating string
		for {
			tokenType := tokenizer.Next()
			if err := tokenizer.Err(); tokenType == html.ErrorToken && err == io.EOF {
				break
			}

			if tokenType == html.SelfClosingTagToken {
				token := tokenizer.Token()
				if "img" == token.Data && strings.TrimSpace(token.Data) != "" {
					tokenType = tokenizer.Next()
					Name = tokenizer.Token().Data
					for _, arr := range token.Attr {
						if arr.Key == "src" {
							UrlImage = arr.Val
							break
						}
					}
				}
			}

			if tokenType == html.StartTagToken {
				token := tokenizer.Token()
				if "strong" == token.Data && strings.TrimSpace(token.Data) != "" {
					tokenType = tokenizer.Next()
					if tokenType == html.TextToken {
						Rating = tokenizer.Token().Data
					}
				}

				if "a" == token.Data && strings.TrimSpace(token.Data) != "" {
					tokenType = tokenizer.Next()
					Name = tokenizer.Token().Data
					for _, arr := range token.Attr {
						if arr.Key == "href" {
							UrlMovie = fmt.Sprintf("https://imdb.com/%s", arr.Val)
							break
						}
					}
				}
			}
		}

		if UrlMovie != "" {
			ch <- models.MovieInfo{Name: Name, UrlMovie: UrlMovie, UrlImage: UrlImage, Rating: Rating}
		}
	}

	close(ch)
}

func SaveToDB(ch chan models.MovieInfo, doneChannel chan bool) {
	for {
		res, ok := <-ch
		if !ok {
			close(doneChannel)
			break
		}

		fmt.Println(res.UrlMovie)
		rawData := GetRawHTML(res.UrlMovie)
		doc, err := html.Parse(strings.NewReader(rawData))
		if err != nil {
			fmt.Println(err.Error())
		}
		allElementData := GetAllRawElementDetails([]string{}, doc)
		if len(allElementData) == 1 {
			movieDetail := GetData(allElementData[0])
			fmt.Println(movieDetail)

			moVieEntity := entities.Movie{
				Url:           movieDetail.Url,
				Name:          movieDetail.Name,
				Image:         movieDetail.Image,
				Description:   movieDetail.Description,
				Keywords:      movieDetail.Keywords,
				DatePublished: movieDetail.DatePublished,
				AggregateRating: entities.AggregateRating{
					Type:        movieDetail.AggregateRating.Type,
					RatingCount: movieDetail.AggregateRating.RatingCount,
					BestRating:  movieDetail.AggregateRating.BestRating,
					WorstRating: movieDetail.AggregateRating.WorstRating,
					RatingValue: movieDetail.AggregateRating.RatingValue,
				},
				Genre: func() (result []entities.Genre) {
					for _, v := range movieDetail.Genre {
						result = append(result, entities.Genre{Genre: v})
					}
					return
				}(),
				Actors: func() (result []entities.Actor) {
					for _, v := range movieDetail.Actors {
						result = append(result, entities.Actor{
							Type: v.Type,
							Url:  v.Url,
							Name: v.Name,
						})
					}
					return
				}(),
				Director: func() (result []entities.Director) {
					for _, v := range movieDetail.Directors {
						result = append(result, entities.Director{
							Type: v.Type,
							Url:  v.Url,
							Name: v.Name,
						})
					}
					return
				}(),
				Creator: func() (result []entities.Creator) {
					for _, v := range movieDetail.Creator {
						result = append(result, entities.Creator{
							Type: v.Type,
							Url:  v.Url,
							Name: v.Name,
						})
					}
					return
				}(),
			}

			resp := DBConn().Create(&moVieEntity)
			fmt.Println(resp)
		}
	}
}
