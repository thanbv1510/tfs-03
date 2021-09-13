package rabbitmq

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"rabbit-demo/models"
	"rabbit-demo/repository"
	"rabbit-demo/sendmail"
)

type Consumer struct {
	ctx        context.Context
	channel    *amqp.Channel
	queue      string
	exchange   string
	exchType   string
	bindingKey string
	db         *sql.DB
}

func NewConsumer(ctx context.Context, channel *amqp.Channel, exchange, exchType, bindingKey, queue string, db *sql.DB) *Consumer {
	return &Consumer{
		ctx:        ctx,
		channel:    channel,
		exchange:   exchange,
		exchType:   exchType,
		bindingKey: bindingKey,
		queue:      queue,
		db:         db,
	}
}

func (consumer *Consumer) declare() error {
	// declare exchange
	if err := consumer.channel.ExchangeDeclare(
		consumer.exchange,
		consumer.exchType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	// declare queue
	queue, err := consumer.channel.QueueDeclare(
		consumer.queue,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil
	}

	if err = consumer.channel.QueueBind(
		queue.Name, consumer.bindingKey, consumer.exchange, false, nil); err != nil {
		return err
	}
	return nil
}

func (consumer *Consumer) Start() {
	consumer.declare()

	delivery, err := consumer.channel.Consume(
		consumer.queue,
		"",
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		return
	}

	for {
		select {
		case mess := <-delivery:
			mess.Ack(false)
			var order models.Order
			_ = json.Unmarshal(mess.Body, &order)
			fmt.Println(order)
			_, err := sendmail.SendEmailThankYou(order.Email, order.Email)
			if err != nil {
				fmt.Println("error: ", err)
				continue
			}

			orderRepository := repository.NewOrderRepository(consumer.db)
			orderRepository.UpdateIsSendThanksEmail(order.ID)

		}
	}
}

func (consumer *Consumer) Close() error {
	return consumer.channel.Close()
}
