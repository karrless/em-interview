package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/karrless/em-interview/internal/models"
	"github.com/karrless/em-interview/pkg/db/postgres"
	"github.com/lib/pq"
)

type SongRepository struct {
	*postgres.DB
}

func NewSongRepository(db *postgres.DB) *SongRepository {
	return &SongRepository{db}
}

// TODO переписать на squirrel
func (r *SongRepository) CreateSong(song *models.Song) (int64, error) {
	query := `INSERT INTO public.songs ("group", song, release_date, "text", link) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	songReleaseDate, err := time.Parse("02.01.2006", song.ReleaseDate)
	if err != nil {
		return 0, err
	}
	var result int64
	err = r.DB.QueryRow(query, song.Group, song.Title, songReleaseDate, song.Text, song.Link).Scan(&result)
	return result, err
}

func (r *SongRepository) GetSong(id int64) (*models.Song, error) {
	query := `SELECT id, "group", song, release_date, "text", link FROM public.songs WHERE id = $1;`
	var song models.Song
	var songReleaseDate time.Time
	err := r.DB.QueryRow(query, id).Scan(&song.ID, &song.Group, &song.Title, &songReleaseDate, &song.Text, &song.Link)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	song.ReleaseDate = songReleaseDate.Format("02.01.2006")
	return &song, nil
}

func (r *SongRepository) DeleteSong(id int64) error {
	query := `DELETE FROM public.songs WHERE id = $1;`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *SongRepository) UpdateSong(song *models.Song) error {
	query := `UPDATE public.songs SET "group" = $1, song = $2, release_date = $3, "text" = $4, link = $5 WHERE id = $6;`
	songReleaseDate, err := time.Parse("02.01.2006", song.ReleaseDate)
	if err != nil {
		return err
	}
	_, err = r.DB.Exec(query, song.Group, song.Title, songReleaseDate, song.Text, song.Link, song.ID)
	return err
}

func (r *SongRepository) GetSongs(filter *models.SongsFilter) ([]*models.Song, error) {
	query := ""
	var args []interface{}
	count := 1
	isParams := false
	if filter.Group != nil && len(*filter.Group) > 0 {
		isParams = true
		query += fmt.Sprintf(`"group" = ANY($%d) `, count)
		count++

		args = append(args, pq.Array(*filter.Group))
	}
	if filter.Title != nil && len(*filter.Title) > 0 {
		isParams = true
		if count != 1 {
			query += `AND `
		}
		query += fmt.Sprintf(`song = ANY($%d) `, count)
		count++
		args = append(args, pq.Array(*filter.Title))
	}
	if filter.ReleaseDate != nil && len(*filter.ReleaseDate) > 0 {
		isParams = true
		if count != 1 {
			query += `AND `
		}
		query += fmt.Sprintf(`release_date = ANY($%d) `, count)
		count++
		args = append(args, filter.ReleaseDate)
	}
	if filter.Before != nil {
		isParams = true
		if count != 1 {
			query += `AND `
		}
		query += fmt.Sprintf(`release_date < $%d `, count)
		count++
		timeBefore, err := time.Parse("02.01.2006", *filter.Before)
		if err != nil {
			return nil, err
		}
		args = append(args, timeBefore)
	}
	if filter.After != nil {
		isParams = true
		if count != 1 {
			query += `AND `
		}
		query += fmt.Sprintf(`release_date > $%d `, count)
		count++
		timeAfter, err := time.Parse("02.01.2006", *filter.After)
		if err != nil {
			return nil, err
		}
		args = append(args, timeAfter)
	}

	query += `ORDER BY id ASC `

	if filter.Limit != nil {
		query += fmt.Sprintf(`LIMIT $%d `, count)
		count++
		args = append(args, filter.Limit)
	}
	if filter.Offset != nil {
		query += fmt.Sprintf(`OFFSET $%d`, count)
		count++
		args = append(args, filter.Offset)
	}

	query += `;`
	if isParams {
		query = "WHERE " + query
	}
	query = `SELECT id, "group", song, release_date, "text", link FROM public.songs ` + query

	rows, err := r.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}

	var songs []*models.Song
	for rows.Next() {
		var song models.Song
		var songReleaseDate time.Time
		err = rows.Scan(&song.ID, &song.Group, &song.Title, &songReleaseDate, &song.Text, &song.Link)
		if err != nil {
			return nil, err
		}
		song.ReleaseDate = songReleaseDate.Format("02.01.2006")

		songs = append(songs, &song)

	}
	return songs, nil
}
