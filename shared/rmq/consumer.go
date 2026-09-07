package rmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Dispatcher interface {
	Dispatch(ctx context.Context)
}

type Consumer struct {
	conn       *amqp.Connection
	channel    *amqp.Channel
	deliveries <-chan amqp.Delivery

	exchangeName string
	kind         string
	queueName    string
	bindKey      string
}

func NewConsumer(
	conn *amqp.Connection,
	exchangeName string,
	kind string,
	queueName string,
	bindKey string,
) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		_ = ch.Close()
		return nil, err
	}

	args := amqp.Table{
		"x-queue-type": "qurom",
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		_ = ch.Close()
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		bindKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		_ = ch.Close()
		return nil, err
	}

	deliveries, err := ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		_ = ch.Close()
		return nil, err
	}

	return &Consumer{
		conn:       conn,
		channel:    ch,
		deliveries: deliveries,

		exchangeName: exchangeName,
		kind:         kind,
		queueName:    queueName,
		bindKey:      bindKey,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) {
	go func() {
		defer func() {
			_ = c.channel.Close()
		}()

		for {
			select {
			case <-ctx.Done():
				return

			case _, ok := <-c.deliveries:
				if !ok {
					return
				}
			}
		}
	}()
}
