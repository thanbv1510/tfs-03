package repository

import (
	"context"
	"database/sql"
	"fmt"
	"rabbit-demo/models"
	"strings"
	"time"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return OrderRepository{db: db}
}

func (orderRepository *OrderRepository) UpdateIsSendThanksEmail(ids ...int) error {
	fmt.Printf("%T %v\n", ids, ids)
	valueStrs := strings.Repeat("?,", len(ids))
	query := fmt.Sprintf("UPDATE orders AS o SET o.is_send_thanks_email = 1 WHERE o.id IN (%s)", valueStrs[:len(valueStrs)-1])
	fmt.Println(query)
	ct, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)

	var valueArgs []interface{}
	for _, v := range ids {
		valueArgs = append(valueArgs, v)
	}

	_, err := orderRepository.db.ExecContext(ct, query, valueArgs...)
	defer cancelFunc()

	if err != nil {
		// log
		fmt.Println(err)
		return err
	}

	return nil
}

func (orderRepository *OrderRepository) GetUnSendEmailsWithLimit(limit int) []models.Order {
	query := "SELECT o.ID, o.EMAIL FROM orders AS o WHERE is_send_thanks_email = 0 LIMIT ?"
	ct, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()

	results, err := orderRepository.db.QueryContext(ct, query, limit)
	if err != nil {
		fmt.Println(err)
		return []models.Order{}
	}

	orders := make([]models.Order, 0)
	for results.Next() {
		var order models.Order
		err := results.Scan(&order.ID, &order.Email)
		if err != nil {
			continue
		}
		orders = append(orders, order)
	}
	return orders
}

func (orderRepository *OrderRepository) CreateOrderTable() error {
	isExistOrderTable := orderRepository.isExistOrderTable()
	if isExistOrderTable {
		fmt.Println("Order Table exist!")
		return nil
	}

	fmt.Println("Creating table...")
	query := `CREATE TABLE orders (
    			id    BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    			username VARCHAR(100) NOT NULL UNIQUE,
    			email VARCHAR(100) NOT NULL UNIQUE,
    			is_send_thanks_email TINYINT(1) DEFAULT 0
            	# More fields ...
			);`

	ct, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	_, err := orderRepository.db.ExecContext(ct, query)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (orderRepository *OrderRepository) isExistOrderTable() bool {
	query := "SELECT COUNT(1) FROM orders"
	ct, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := orderRepository.db.QueryContext(ct, query)
	defer cancelFunc()

	if err != nil {
		fmt.Println("Table not exist... Please create table")
		return false
	}

	return true
}
