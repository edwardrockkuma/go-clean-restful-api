package server

import (
	"edward/cleanarc/pkg/api/handler"
	middleware "edward/cleanarc/pkg/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp(userHandler *handler.UserHandler) *ServerHttp {
	engine := gin.New()

	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	engine.POST("/login", middleware.LoginHandler)

	// Auth middleware
	api := engine.Group("/api", middleware.AuthorizationMiddleware)

}
