package rmq

import (
	"fmt"
	"net/url"

	"ops/shared/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConnection(cfg *config.Env) (*amqp.Connection, error) {
	vhost := url.PathEscape(cfg.MQ.VHost)

	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		cfg.MQ.User,
		cfg.MQ.Password,
		cfg.MQ.Host,
		cfg.MQ.Port,
		vhost,
	)

	conn, err := amqp.Dial(connAddr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
