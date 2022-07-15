package v1

import (
	"web-socket/internal/usecase"
	"web-socket/pkg/logger"

	_ "web-socket/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Swagger spec:
// @title			Go Clean chat app with websocket API
// @description		Using chat app as experiment
// @version 		1.0
// @host			localhost:90
// @BasePath		/v1
func NewRouter(handler *gin.Engine, t *usecase.MessageUseCase, l logger.Interface) {
	// Options ...
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)

	h := handler.Group("/v1")
	{
		newMessageRoutes(h, t, l)
	}
	handler.GET("/swagger/*any", swaggerHandler)
}
