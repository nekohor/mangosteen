package app

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/nekohor/mangosteen/internal/app/config"
	"github.com/nekohor/mangosteen/internal/app/middleware"
	"github.com/nekohor/mangosteen/internal/app/router"
	_ "github.com/nekohor/mangosteen/docs/swagger"
	"github.com/nekohor/mangosteen/pkg/logger/zap/logger"
)

func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	prefixes := r.Prefixes()

	// Trace ID
	app.Use(middleware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Copy body
	app.Use(middleware.CopyBodyMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Access logger
	app.Use(middleware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Recover
	app.Use(middleware.RecoveryMiddleware())

	// CORS
	if config.C.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	// Router register
	err := r.Register(app); if err != nil {
		logger.Error("Register Routers Failed.")
	}

	// Swagger
	if config.C.Swagger {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Website
	if dir := config.C.WWW; dir != "" {
		app.Use(middleware.WWWMiddleware(dir, middleware.AllowPathPrefixSkipper(prefixes...)))
	}
	return app
}