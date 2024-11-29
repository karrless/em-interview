package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/karrless/em-interview/internal/service"
	"github.com/karrless/em-interview/internal/transport/rest/controllers"
)

func SongsRoutes(ctx *context.Context, r *gin.RouterGroup, songsService *service.SongsService) {

	songsController := controllers.NewSongsController(ctx, songsService)

	songsGroup := r.Group("/songs")
	{
		songsGroup.POST("", songsController.CreateSong)
		songsGroup.GET("/:id", songsController.GetSong)
		songsGroup.DELETE("/:id", songsController.DeleteSong)
		songsGroup.PUT("/:id", songsController.UpdateSong)
		songsGroup.GET("", songsController.GetSongs)
	}

}
