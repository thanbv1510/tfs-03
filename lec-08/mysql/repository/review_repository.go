package repository

import (
	"context"
	"csv-import-mysql/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type ReviewStore struct {
	db *sql.DB
}

func InsertBatchReview(reviews []models.Review, db *sql.DB) error {
	defer db.Close()
	var valueStrs []string
	var valueArgs []interface{}
	for _, v := range reviews {
		valueStrs = append(valueStrs, "(?,	?,	?)")
		valueArgs = append(valueArgs, v.Type, v.Title, v.Body)
	}
	query := fmt.Sprintf("INSERT INTO review(type, title, body) VALUES %s", strings.Join(valueStrs, ","))

	ct, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	res, err := db.ExecContext(ct, query, valueArgs...)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func CreateReviewTable(db *sql.DB) error {
	query := `create table review (
    			id    bigint auto_increment primary key,
    			type  varchar(1) charset utf8 null,
    			title varchar(255) charset utf8 null,
    			body  mediumtext charset utf8 null
			);`

	ct, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	res, err := db.ExecContext(ct, query)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func FindByBodyReview(keyword string, sqlDB *sql.DB) []*models.Review {
	query := fmt.Sprintf("SELECT r.type, r.title, r.body FROM review r WHERE r.body LIKE '%%%s%%'", keyword)
	fmt.Println(query)
	ct, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFunc()
	result, err := sqlDB.QueryContext(ct, query)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var reviews []*models.Review
	for result.Next() {
		var review models.Review
		err := result.Scan(&review.Type, &review.Title, &review.Body)
		if err != nil {
			continue
		}

		reviews = append(reviews, &review)
	}

	return reviews
}
