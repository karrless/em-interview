package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karrless/em-interview/internal/models"
	"github.com/karrless/em-interview/pkg/logger"
	"go.uber.org/zap"
)

type SongsService interface {
	CreateSong(song *models.Song) (*models.Song, error)
	GetSong(id int64) (*models.Song, error)
	DeleteSong(id int64) error
	UpdateSong(song *models.Song) error
	GetSongs(filter *models.SongsFilter) ([]*models.Song, error)
}

type SongsController struct {
	service SongsService
	ctx     *context.Context
}

func NewSongsController(ctx *context.Context, service SongsService) *SongsController {
	return &SongsController{service: service, ctx: ctx}
}

// @Summary Get song by id
// Tags Songs
// @Produce  json
// @Param id path int true "Song ID"
// @Success 200 {object} models.Song
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Failure 500 {object} error "Internal server error"
// @Router /songs/{id} [get]
func (sc *SongsController) GetSong(c *gin.Context) {
	debugLogger := logger.GetLoggerFromCtx(*sc.ctx)
	id := c.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, fmt.Errorf("ID must be a number"))
		return
	}
	song, err := sc.service.GetSong(id64)
	if err != nil {
		debugLogger.Debug("GetSong", zap.Error(err))
		c.JSON(500, fmt.Errorf("Internal server error"))
		return
	}
	if song == nil {
		c.JSON(404, fmt.Errorf("Song with ID %d not found", id64))
		return
	}
	c.JSON(200, song)
}

// @Summary Delete song by id
// Tags Songs
// @Produce  json
// @Param id path int true "Song ID"
// @Success 200 {object} nil
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal server error"
// @Router /songs/{id} [delete]
func (sc *SongsController) DeleteSong(c *gin.Context) {
	debugLogger := logger.GetLoggerFromCtx(*sc.ctx)
	id := c.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, fmt.Errorf("ID must be a number"))
		return
	}
	err = sc.service.DeleteSong(id64)
	if err != nil {
		debugLogger.Debug("DeleteSong", zap.Error(err))
		c.JSON(500, fmt.Errorf("Internal server error"))
		return
	}
	c.JSON(200, nil)
}

// @Summary Update song by id
// Tags Songs
// @Produce  json
// @Param id path int true "Song ID"
// @Param song body models.Song true "Song"
// @Success 200 {object} nil
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal server error"
// @Router /songs/{id} [put]
func (sc *SongsController) UpdateSong(c *gin.Context) {
	debugLogger := logger.GetLoggerFromCtx(*sc.ctx)
	id := c.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, fmt.Errorf("ID must be a number"))
		return
	}
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(400, err)
		return
	}
	song.ID = id64
	err = sc.service.UpdateSong(&song)
	if err != nil {
		debugLogger.Debug("UpdateSong", zap.Error(err))
		c.JSON(500, fmt.Errorf("Internal server error"))
		return
	}
	c.JSON(200, nil)
}

// @Summary Create song
// Tags Songs
// @Produce  json
// @Param title body string true "Song title"
// @Param group body string true "Song title"
// @Success 200 {object} models.Song "Created song"
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal server error"
// @Router /songs [post]
func (sc *SongsController) CreateSong(c *gin.Context) {
	debugLogger := logger.GetLoggerFromCtx(*sc.ctx)
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(400, err)
		return
	}
	createdSong, err := sc.service.CreateSong(&song)
	if err != nil {
		debugLogger.Debug("CreateSong", zap.Error(err))
		c.JSON(500, fmt.Errorf("Internal server error"))
		return
	}
	c.JSON(201, createdSong)
}

// @Summary Get songs
// Tags Songs
// @Produce  json
// @Param group query string false "Song group"
// @Param title query string false "Song title"
// @Param release_date query string false "Song release date"
// @Param before query string false "Song release date before"
// @Param after query string false "Song release date after"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} []models.Song
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal server error"
// @Router /songs [get]
func (sc *SongsController) GetSongs(c *gin.Context) {
	debugLogger := logger.GetLoggerFromCtx(*sc.ctx)
	var filter models.SongsFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(400, err)
		return
	}
	songs, err := sc.service.GetSongs(&filter)
	if err != nil {
		debugLogger.Debug("GetSongs", zap.Error(err))
		c.JSON(500, fmt.Errorf("Internal server error"))
		return
	}
	c.JSON(200, songs)
}
