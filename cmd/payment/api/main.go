package main

import (
	"fmt"
	"os"

	paymentusecase "ops/payment/application/usecase"
	paymenthandler "ops/payment/delivery/rest/handler"
	"ops/payment/delivery/rest/router"
	paymentdb "ops/payment/infra/database"
	"ops/shared/config"
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

	processPaymentUseCase := paymentusecase.NewProcessPaymentUseCase()
	h := paymenthandler.NewHandler(processPaymentUseCase)
	_ = h

	router.NewRoutes(appServer)

	port := env.HTTP.PaymentPort
	if port == 0 {
		port = 8083
	}

	if err := appServer.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("error run payment server: %v\n", err)
		os.Exit(1)
	}
}
