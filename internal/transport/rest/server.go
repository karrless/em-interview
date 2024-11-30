package rest

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/karrless/em-interview/internal/service"
	"github.com/karrless/em-interview/internal/transport/rest/middlewares"
	"github.com/karrless/em-interview/internal/transport/rest/routes"
	"github.com/karrless/em-interview/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"github.com/karrless/em-interview/docs"
)

type ServerConfig struct {
	Host string `env:"SERVER_HOST" env-default:"localhost"`
	Port string `env:"SERVER_PORT" env-default:"8080"`
}

type Server struct {
	r      *gin.Engine
	config ServerConfig
}

func New(ctx *context.Context, cfg ServerConfig, SongsService *service.SongsService, debug bool) *Server {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(middlewares.WithLogger(ctx), gin.Recovery())
	r.SetTrustedProxies([]string{"127.0.0.1", cfg.Host})
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = cfg.Host + ":" + cfg.Port
	docs.SwaggerInfo.Title = "MusicLibrary API"
	docs.SwaggerInfo.Description = "API for music library website"
	docs.SwaggerInfo.Version = "0.1.0"

	api := r.Group("/api")
	v1 := api.Group("/v1")

	routes.HealthCheckRoutes(v1)
	routes.SongsRoutes(ctx, v1, SongsService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return &Server{r: r, config: cfg}
}

func (s *Server) Run(ctx *context.Context) error {
	logger.GetLoggerFromCtx(*ctx).Info("Server started", zap.String("host", s.config.Host), zap.String("port", s.config.Port))
	return s.r.Run(s.config.Host + ":" + s.config.Port)
}
