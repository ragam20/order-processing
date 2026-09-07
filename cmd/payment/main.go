package main

import (
	"fmt"
	"net/http"
	"os"

	paymentusecase "ops/payment/application/usecase"
	paymenthandler "ops/payment/delivery/rest/handler"
	paymentdb "ops/payment/infra/database"
	"ops/shared/config"
	"ops/shared/server"

	"github.com/gin-gonic/gin"
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

	processPaymentUseCase := paymentusecase.NewProcessPaymentUseCase()
	h := paymenthandler.NewHandler(processPaymentUseCase)
	_ = h

	appServer := server.New(env)
	api := appServer.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "payment",
			"status": "healthy",
		})
	})

	port := env.HTTP.PaymentPort
	if port == 0 {
		port = env.HTTP.Port
	}
	if port == 0 {
		port = 8083
	}

	if err := appServer.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("error run payment server: %v\n", err)
		os.Exit(1)
	}
}
