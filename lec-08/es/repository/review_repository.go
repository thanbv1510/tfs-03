package repository

import (
	"context"
	"csv-import-es/models"
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
