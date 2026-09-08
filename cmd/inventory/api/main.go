package main

import (
	"fmt"
	"os"

	inventoryusecase "ops/inventory/application/usecase"
	inventoryhandler "ops/inventory/delivery/rest/handler"
	"ops/inventory/delivery/rest/router"
	inventorydb "ops/inventory/infra/database"
	"ops/shared/config"
	"ops/shared/server"
)

func main() {
	env := config.LoadEnv()

	db, err := inventorydb.NewPostgres(env)
	if err != nil {
		panic(err)
	}

	if err := inventorydb.Migrate(db); err != nil {
		panic(err)
	}

	appServer := server.New(env)

	repository := inventorydb.NewRepository(db)

	getProductsUseCase := inventoryusecase.NewGetProductsUseCase(repository)
	getProductByIDUseCase := inventoryusecase.NewGetProductByIDUseCase(repository)
	checkAvailabilityUseCase := inventoryusecase.NewCheckAvailability(repository)

	h := inventoryhandler.NewHandler(getProductsUseCase, getProductByIDUseCase, checkAvailabilityUseCase)

	router.NewRoutes(appServer, h)

	port := env.HTTP.InventoryPort
	if port == 0 {
		port = 8081
	}

	if err := appServer.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("error run inventory server: %v\n", err)
		os.Exit(1)
	}
}
