package main

import (
	"fmt"
	"os"

	paymentusecase "ops/payment/application/usecase"
	paymenthandler "ops/payment/delivery/rest/handler"
	"ops/payment/delivery/rest/router"
	paymentdb "ops/payment/infra/database"
	"ops/payment/infra/messaging"
	"ops/payment/infra/provider/midtrans"
	"ops/shared/config"
	"ops/shared/rmq"
	"ops/shared/server"
)

func main() {
	env := config.LoadEnv()

	db, err := paymentdb.NewPostgres(env)
	if err != nil {
		panic(err)
	}

	if err := paymentdb.Migrate(db); err != nil {
		panic(err)
	}

	appServer := server.New(env)

	repository := paymentdb.NewRepository(db)
	provider := midtrans.New(env)
	mqConnection, err := rmq.NewConnection(env)
	if err != nil {
		panic(err)
	}
	defer mqConnection.Close()

	rmqPublisher := rmq.NewPublisher(mqConnection)
	paymentEventPublisher := messaging.NewPaymentEventPublisher(rmqPublisher)
	processStatusPaymentUseCase := paymentusecase.NewProcessStatusPaymentUseCase(repository, paymentEventPublisher)
	createPaymentUseCase := paymentusecase.NewCreatePaymentUseCase(repository, provider, paymentEventPublisher)
	getStatusPaymentUseCase := paymentusecase.NewGetPaymentStatus(provider)

	h := paymenthandler.NewHandler(createPaymentUseCase, getStatusPaymentUseCase)
	wbHandler := paymenthandler.NewWebHookHandler(processStatusPaymentUseCase, provider)

	router.NewRoutes(appServer, h)
	router.NewWebhooks(appServer, wbHandler)

	port := env.HTTP.PaymentPort
	if port == 0 {
		port = 8083
	}

	if err := appServer.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("error run payment server: %v\n", err)
		os.Exit(1)
	}
}
