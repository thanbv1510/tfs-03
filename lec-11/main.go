package main

import (
	"context"
	"fmt"
	"rabbit-demo/database"
	"rabbit-demo/rabbitmq"
	"rabbit-demo/repository"
	"rabbit-demo/scheduler"
)

func main() {
	db, _ := database.DBConn()
	orderRepository := repository.NewOrderRepository(db)

	err := orderRepository.CreateOrderTable()
	if err != nil {
		fmt.Println("Cannot create orders table.Please try again :(")
		return
	}

	ctx := context.Background()
	rmqURI := "amqp://guest:guest@localhost:5672"
	exch := "order"
	exchType := "direct"
	queue := "order_processor"
	routingKey := ""

	// connect to rabbitmq
	rmq := rabbitmq.NewRMQ(rmqURI)

	// create 1 channel for producer
	pCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cann't get channel: ", err)
		return
	}
	// create 1 channel for consumer
	cCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cann't get channel: ", err)
		return
	}

	producer := rabbitmq.NewProducer(ctx, pCh, exch, exchType, routingKey)
	consumer := rabbitmq.NewConsumer(ctx, cCh, exch, exchType, routingKey, queue, db)

	scheduler := scheduler.NewScheduler(ctx, db, producer)
	scheduler.Start()
	consumer.Start()

}
