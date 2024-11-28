package repository

import (
	"context"
	"testing"

	"github.com/karrless/em-interview/internal/config"
	"github.com/karrless/em-interview/internal/models"
	"github.com/karrless/em-interview/pkg/db/postgres"
)

func TestCreateSong(t *testing.T) {
	cfg := config.New("./../../.env")
	ctx := context.Background()
	db, err := postgres.New(&ctx, cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewSongRepository(db)
	resultID, err := repo.CreateSong(&models.Song{
		Group:       "group",
		Title:       "title",
		ReleaseDate: "28.12.2003",
		Text:        "text",
		Link:        "link",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resultID)
}

func TestGetSong(t *testing.T) {
	cfg := config.New("./../../.env")
	ctx := context.Background()
	db, err := postgres.New(&ctx, cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewSongRepository(db)
	song, err := repo.GetSong(3)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(song)
	song, err = repo.GetSong(1)
	if err != nil || song != nil {
		t.Fatal(err)
	}
	t.Log(song)
}

func TestUpdateSong(t *testing.T) {
	cfg := config.New("./../../.env")
	ctx := context.Background()
	db, err := postgres.New(&ctx, cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewSongRepository(db)
	err = repo.UpdateSong(&models.Song{
		ID:          2,
		Group:       "test",
		Title:       "title",
		ReleaseDate: "28.12.2003",
		Text:        "text",
		Link:        "link",
	})
	if err != nil {
		t.Fatal(err)
	}

	song, err := repo.GetSong(2)
	if err != nil {
		t.Fatal(err)
	}
	if song.Group != "test" {
		t.Fatal("wrong group")
	}
}

func TestDeleteSong(t *testing.T) {
	cfg := config.New("./../../.env")
	ctx := context.Background()
	db, err := postgres.New(&ctx, cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewSongRepository(db)
	err = repo.DeleteSong(3)
	if err != nil {
		t.Fatal(err)
	}
	song, err := repo.GetSong(3)
	if err != nil || song != nil {
		t.Fatal(err)
	}
	t.Log(song)
}

func TestGetSongs(t *testing.T) {
	cfg := config.New("./../../.env")
	ctx := context.Background()
	db, err := postgres.New(&ctx, cfg.PostgresConfig)
	if err != nil {
		t.Fatal(err)
	}
	filter := models.SongsFilter{}
	repo := NewSongRepository(db)
	songs, err := repo.GetSongs(&filter)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", songs)
	date := "28.12.2003"
	filter.After = &date
	songs, err = repo.GetSongs(&filter)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(songs)
	filter.After = nil
	filter.Before = &date
	songs, err = repo.GetSongs(&filter)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(songs)
	filter.Before = nil
	num := 1
	filter.Limit = &num
	songs, err = repo.GetSongs(&filter)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(songs)
	filter.Limit = nil
	filter.Offset = &num
	songs, err = repo.GetSongs(&filter)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(songs)
	filter.Offset = nil
	filter.Group = &[]string{"test"}
	songs, err = repo.GetSongs(&filter)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(songs)
	filter.Group = nil
	filter.Title = &[]string{"title"}
	songs, err = repo.GetSongs(&filter)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(songs)
	filter.Group = &[]string{"test"}
	filter.Title = &[]string{"title"}
	songs, err = repo.GetSongs(&filter)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(songs)
}
