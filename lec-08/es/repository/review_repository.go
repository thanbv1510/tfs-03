package repository

import (
	"context"
	"csv-import-es/models"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

const indexName = "reviews"

func InsertBulkReview(reviews []models.Review, client *elastic.Client) {
	bulk := client.Bulk()

	for _, review := range reviews {
		req := elastic.NewBulkIndexRequest()
		req.OpType("index")
		req.Index(indexName)
		req.Doc(review)

		bulk = bulk.Add(req)
	}

	_, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}

func FindByBodyReview(keyword string, client *elastic.Client) []*models.Review {
	if len(keyword) == 0 {
		return nil
	}

	query := elastic.NewSearchSource()
	query.Query(elastic.NewMultiMatchQuery(keyword, "Body").Type("phrase_prefix")).Size(1000)

	searchService := client.Search().Index(indexName).SearchSource(query)

	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var reviews []*models.Review
	for _, v := range searchResult.Hits.Hits {
		var review models.Review
		err := json.Unmarshal(v.Source, &review)
		if err != nil {
			continue
		}

		reviews = append(reviews, &review)
	}

	return reviews
}
