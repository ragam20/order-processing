package main

import (
	"fmt"
	"net/http"
	"os"

	notificationusecase "ops/notification/application/usecase"
	notificationhandler "ops/notification/delivery/rest/handler"
	notificationdb "ops/notification/infra/database"
	"ops/shared/config"
	"ops/shared/server"

	"github.com/gin-gonic/gin"
)

func main() {
	env := config.LoadEnv()

	db, err := notificationdb.NewPostgres(env)
	if err != nil {
		panic(err)
	}

	if err := notificationdb.Migrate(db); err != nil {
		panic(err)
	}

	repository := notificationdb.NewRepository(db)
	getNotificationsUseCase := notificationusecase.NewGetNotificationsUseCase(repository)
	createNotificationUseCase := notificationusecase.NewCreateUseCase(repository)
	h := notificationhandler.NewHandler(getNotificationsUseCase, createNotificationUseCase)
	_ = h

	appServer := server.New(env)
	api := appServer.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "notification",
			"status": "healthy",
		})
	})

	port := env.HTTP.NotificationPort
	if port == 0 {
		port = env.HTTP.Port
	}
	if port == 0 {
		port = 8082
	}

	if err := appServer.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("error run notification server: %v\n", err)
		os.Exit(1)
	}
}
