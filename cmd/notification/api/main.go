package main

import (
	"fmt"
	"os"

	notificationusecase "ops/notification/application/usecase"
	notificationhandler "ops/notification/delivery/rest/handler"
	"ops/notification/delivery/rest/router"
	notificationdb "ops/notification/infra/database"
	"ops/shared/config"
	"ops/shared/server"
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

	appServer := server.New(env)

	repository := notificationdb.NewRepository(db)
	getNotificationsUseCase := notificationusecase.NewGetNotificationsUseCase(repository)
	createNotificationUseCase := notificationusecase.NewCreateUseCase(repository)
	h := notificationhandler.NewHandler(getNotificationsUseCase, createNotificationUseCase)
	_ = h

	router.NewRoutes(appServer)

	port := env.HTTP.NotificationPort
	if port == 0 {
		port = 8082
	}

	if err := appServer.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("error run notification server: %v\n", err)
		os.Exit(1)
	}
}
