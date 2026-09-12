package main

import (
	"fmt"
	"os"

	orderusecase "ops/order/application/usecase"
	orderhandler "ops/order/delivery/rest/handler"
	"ops/order/delivery/rest/router"
	orderdb "ops/order/infra/database"
	"ops/order/infra/inventory"
	"ops/order/infra/payment"
	"ops/shared/config"
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

	appServer := server.New(env)

	repo := orderdb.NewRepository(db)

	// internal service
	inventorySvc := inventory.New(env)
	paymentSvc := payment.New(env)

	cancelOrderUC := orderusecase.NewCancelOrderUseCase(repo)
	getDetailOrderUC := orderusecase.NewGetDetailOrderUseCase(repo)
	checkoutUC := orderusecase.NewCheckoutUseCase(repo, inventorySvc, paymentSvc)

	h := orderhandler.NewHandler(cancelOrderUC, getDetailOrderUC, checkoutUC)

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
