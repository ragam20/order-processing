package rmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumerConfig struct {
	ExchangeName string
	ExchangeType string
	QueueName    string
	BindingKey   string
}

type MessageHandler func(ctx context.Context, message amqp.Delivery) error

type Consumer interface {
	Start(ctx context.Context)
}

type consumer struct {
	conn       *amqp.Connection
	channel    *amqp.Channel
	deliveries <-chan amqp.Delivery
	handler    MessageHandler
}

func NewConsumer(
	conn *amqp.Connection,
	config ConsumerConfig,
	handler MessageHandler,
) (Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		config.ExchangeName,
		config.ExchangeType,
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
		"x-queue-type": "quorum",
	}

	q, err := ch.QueueDeclare(
		config.QueueName,
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
		config.BindingKey,
		config.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		_ = ch.Close()
		return nil, err
	}

	deliveries, err := ch.Consume(
		config.QueueName,
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

	return &consumer{
		conn:       conn,
		channel:    ch,
		deliveries: deliveries,
		handler:    handler,
	}, nil
}

func (c *consumer) Start(ctx context.Context) {
	go func() {
		defer func() {
			_ = c.channel.Close()
		}()

		for {
			select {
			case <-ctx.Done():
				return

			case delivery, ok := <-c.deliveries:
				if !ok {
					return
				}

				if err := c.handler(ctx, delivery); err != nil {
					_ = delivery.Nack(false, false)
					continue
				}

				_ = delivery.Ack(false)
			}
		}
	}()
}
