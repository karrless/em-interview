package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/karrless/em-interview/internal/models"
)

var (
	ErrBadRequest = fmt.Errorf("bad request")
	ErrNoResponse = fmt.Errorf("no response")
)

type ExternalAPIConfig struct {
	ExtarnalAPIURL string `env:"EXTERNAL_API_URL" env-default:"http://localhost:9090"`
}

type ExtarnalAPIRepository struct {
	config *ExternalAPIConfig
}

func NewExtarnalAPIRepository(config *ExternalAPIConfig) *ExtarnalAPIRepository {
	return &ExtarnalAPIRepository{config: config}
}

func (r *ExtarnalAPIRepository) UpdateSongInfo(song *models.Song) error {
	request := fmt.Sprintf("%s/info?song=%s&group=%s", r.config.ExtarnalAPIURL, song.Title, song.Group)
	resp, err := http.Get(request)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return ErrBadRequest
		}
		return ErrNoResponse
	}

	if err := json.NewDecoder(resp.Body).Decode(song); err != nil {
		return err
	}
	return nil
}
