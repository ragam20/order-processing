package main

import (
	"fmt"
	"net/http"
	"os"

	inventoryusecase "ops/inventory/application/usecase"
	inventoryhandler "ops/inventory/delivery/rest/handler"
	inventorydb "ops/inventory/infra/database"
	"ops/shared/config"
	"ops/shared/server"

	"github.com/gin-gonic/gin"
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

	repository := inventorydb.NewRepository(db)
	getProductsUseCase := inventoryusecase.NewGetProductsUseCase(repository)
	getProductByIDUseCase := inventoryusecase.NewGetProductByIDUseCase(repository)
	h := inventoryhandler.NewHandler(getProductsUseCase, getProductByIDUseCase)
	_ = h

	appServer := server.New(env)
	api := appServer.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "inventory",
			"status": "healthy",
		})
	})

	port := env.HTTP.InventoryPort
	if port == 0 {
		port = env.HTTP.Port
	}
	if port == 0 {
		port = 8081
	}

	if err := appServer.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("error run inventory server: %v\n", err)
		os.Exit(1)
	}
}
