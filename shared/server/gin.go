package server

import (
	"time"

	"ops/shared/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(env *config.Env) *gin.Engine {
	appServer := gin.New()

	appServer.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	appServer.Use(gin.Logger(), gin.Recovery())

	return appServer
}
