package main

import (
	"fmt"
	"net/http"
	"os"

	orderusecase "ops/order/application/usecase"
	orderhandler "ops/order/delivery/rest/handler"
	orderdb "ops/order/infra/database"
	"ops/shared/config"
	"ops/shared/server"

	"github.com/gin-gonic/gin"
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

	repo := orderdb.NewRepository(db)

	cancelOrderUseCase := orderusecase.NewCancelOrderUseCase(repo)
	getDetailOrderUseCase := orderusecase.NewGetDetailOrderUseCase(repo)
	h := orderhandler.NewHandler(cancelOrderUseCase, getDetailOrderUseCase)
	_ = h

	appServer := server.New(env)
	api := appServer.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "order",
			"status": "healthy",
		})
	})

	port := env.HTTP.OrderPort
	if port == 0 {
		port = 8080
	}

	if err := appServer.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("error run order server: %v\n", err)
		os.Exit(1)
	}
}
