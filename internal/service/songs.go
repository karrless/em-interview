package service

import "github.com/karrless/em-interview/internal/models"

type SongsRepository interface {
	CreateSong(song *models.Song) (int64, error)
	GetSong(id int64) (*models.Song, error)
	DeleteSong(id int64) error
	UpdateSong(song *models.Song) error
	GetSongs(filter *models.SongsFilter) ([]*models.Song, error)
}

type ExternalAPIRepository interface {
	UpdateSongInfo(song *models.Song) error
}

type SongsService struct {
	songsRepo       SongsRepository
	externalAPIRepo ExternalAPIRepository
}

func NewSongsService(songsRepo SongsRepository, externalAPIRepo ExternalAPIRepository) *SongsService {

	return &SongsService{songsRepo: songsRepo, externalAPIRepo: externalAPIRepo}
}

func (s *SongsService) CreateSong(song *models.Song) (int64, error) {
	if err := s.externalAPIRepo.UpdateSongInfo(song); err != nil {
		return 0, err
	}
	return s.songsRepo.CreateSong(song)
}

func (s *SongsService) GetSong(id int64) (*models.Song, error) {
	return s.songsRepo.GetSong(id)
}

func (s *SongsService) DeleteSong(id int64) error {
	return s.songsRepo.DeleteSong(id)
}

func (s *SongsService) UpdateSong(song *models.Song) error {
	return s.songsRepo.UpdateSong(song)
}

func (s *SongsService) GetSongs(filter *models.SongsFilter) ([]*models.Song, error) {
	return s.songsRepo.GetSongs(filter)
}
