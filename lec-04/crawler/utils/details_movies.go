package utils

import (
	"crawler/models"
	"encoding/json"
	"fmt"
		"golang.org/x/net/html"
	"io"
	"strings"
)

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
						fmt.Println(err)
					}

					return
				}
			}
		}
	}
	return
}
