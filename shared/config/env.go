package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	HTTP struct {
		OrderPort        int `mapstructure:"order_http_port"`
		InventoryPort    int `mapstructure:"inventory_http_port"`
		NotificationPort int `mapstructure:"notification_http_port"`
		PaymentPort      int `mapstructure:"payment_http_port"`

		OrderURL        string `mapstructure:"order_service_url"`
		InventoryURL    string `mapstructure:"inventory_service_url"`
		NotificationURL string `mapstructure:"notification_service_url"`
		PaymentURL      string `mapstructure:"payment_service_url"`
	}
	MQ struct {
		Port     int    `mapstructure:"mq_port"`
		Password string `mapstructure:"mq_password"`
		Host     string `mapstructure:"mq_host"`
		User     string `mapstructure:"mq_user"`
		VHost    string `mapstructure:"mq_vhost"`
	}
	Database struct {
		Host             string `mapstructure:"db_host"`
		User             string `mapstructure:"db_user"`
		Name             string `mapstructure:"db_name"`
		Port             int    `mapstructure:"db_port"`
		Password         string `mapstructure:"db_password"`
		SSLMode          string `mapstructure:"db_sslmode"`
		OrderName        string `mapstructure:"db_order_name"`
		InventoryName    string `mapstructure:"db_inventory_name"`
		NotificationName string `mapstructure:"db_notification_name"`
		PaymentName      string `mapstructure:"db_payment_name"`
	}
	Midtrans struct {
		MerchantID string `mapstructure:"midtrans_merchant_id"`
		ServerKey  string `mapstructure:"midtrans_server_key"`
		SandboxURL string `mapstructure:"midtrans_sandbox_url"`
	}
}

func getKey() []string {
	return []string{
		"order_http_port", "inventory_http_port", "notification_http_port", "payment_http_port",
		"order_service_url", "inventory_service_url", "notification_service_url", "payment_service_url",
		"db_host", "db_user", "db_name", "db_port", "db_password", "db_sslmode",
		"db_order_name", "db_inventory_name", "db_notification_name", "db_payment_name",
		"mq_port", "mq_user", "mq_host", "mq_password", "mq_vhost",
		"midtrans_merchant_id", "midtrans_server_key", "midtrans_sandbox_url",
	}
}

func LoadEnv() *Env {
	v := viper.New()
	v.AutomaticEnv()

	for _, k := range getKey() {
		if err := v.BindEnv(k); err != nil {
			panic(fmt.Sprintf("error bind key: %s", k))
		}
	}

	var env Env
	if err := v.Unmarshal(&env); err != nil {
		panic("error unmarshal!")
	}

	return &env
}
