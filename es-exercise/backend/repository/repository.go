package repository

import (
	"awesomeProject/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

const indexName = "data"

func InsertBulk(reviews []model.Data, client *elastic.Client) {
	bulk := client.Bulk()

	for _, item := range reviews {
		req := elastic.NewBulkIndexRequest()
		req.OpType("index")
		req.Index(indexName)
		req.Doc(item)

		bulk = bulk.Add(req)
	}

	_, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}

func FindData(keyword string, client *elastic.Client) []*model.Data {
	if len(keyword) == 0 {
		return nil
	}

	query := elastic.NewSearchSource()
	query.Query(elastic.NewMultiMatchQuery(keyword, "Text", "FileName", "MEDShortTitle", "SourceCorpus", "Edn", "MS", "MEDData", "Area", "VP", "Genre", "Words").Type("phrase_prefix")).Size(1000)

	searchService := client.Search().Index(indexName).SearchSource(query)

	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var data []*model.Data
	for _, v := range searchResult.Hits.Hits {
		var item model.Data
		err := json.Unmarshal(v.Source, &item)
		if err != nil {
			continue
		}

		data = append(data, &item)
	}

	return data
}
