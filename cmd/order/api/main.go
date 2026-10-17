package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	orderusecase "ops/order/application/usecase"
	evHandler "ops/order/delivery/event"
	orderhandler "ops/order/delivery/rest/handler"
	"ops/order/delivery/rest/router"
	orderdb "ops/order/infra/database"
	"ops/order/infra/inventory"
	"ops/order/infra/payment"
	"ops/shared/config"
	"ops/shared/event"
	"ops/shared/rmq"
	"ops/shared/server"
)

func main() {
	env := config.LoadEnv()

	db, err := orderdb.NewPostgres(env)
	if err != nil {
		panic(err)
	}

	if err := orderdb.Migrate(db); err != nil {
		panic(err)
	}

	rootCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	appServer := server.New(env)

	repo := orderdb.NewRepository(db)

	// rmq
	rmqConfig, err := rmq.NewConnection(env)
	if err != nil {
		panic(err)
	}
	defer rmqConfig.Close()

	// internal service
	inventorySvc := inventory.New(env)
	paymentSvc := payment.New(env)

	cancelOrderUC := orderusecase.NewCancelOrderUseCase(repo)
	getDetailOrderUC := orderusecase.NewGetDetailOrderUseCase(repo)
	getOrdersUC := orderusecase.NewGetOrdersUseCase(repo)
	getOrdersByStatusUC := orderusecase.NewGetOrdersByStatusUseCase(repo)
	checkoutUC := orderusecase.NewCheckoutUseCase(repo, inventorySvc, paymentSvc)
	processOrderStatusUC := orderusecase.NewProcessOrderStatusUseCase(repo)

	h := orderhandler.NewHandler(cancelOrderUC, getDetailOrderUC, getOrdersUC, getOrdersByStatusUC, checkoutUC)
	eh := evHandler.NewEventHandler(processOrderStatusUC)

	processOrderConsumer, err := rmq.NewConsumer(
		rmqConfig,
		rmq.ConsumerConfig{
			ExchangeName: evHandler.PaymentEventsExchange,
			ExchangeType: event.ExchangeTypeTopic,
			QueueName:    evHandler.QueueName,
			BindingKeys: []string{
				evHandler.PaymentPaidRoutingKey,
				evHandler.PaymentFailedRoutingKey,
				evHandler.PaymentExpiredRoutingKey,
			},
		},
		eh.ProcessOrderStatus,
	)
	if err != nil {
		panic(err)
	}
	processOrderConsumer.Start(rootCtx)

	router.NewRoutes(appServer, h)

	port := env.HTTP.OrderPort
	if port == 0 {
		port = 8080
	}

	if err := appServer.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("error run order server: %v\n", err)
		os.Exit(1)
	}
}
