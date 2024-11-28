package models

type Song struct {
	ID          int64  `json:"id" db:"id"`
	Group       string `json:"group" db:"group"`
	Title       string `json:"song" db:"song"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
	Text        string `json:"text" db:"text"`
	Link        string `json:"link" db:"link"`
}

type SongsFilter struct {
	Group       *[]string `json:"group" db:"group"`
	Title       *[]string `json:"song" db:"song"`
	ReleaseDate *[]string `json:"releaseDate" db:"release_date"`
	Before      *string   `json:"before" db:"release_date"`
	After       *string   `json:"after" db:"release_date"`
	Offset      *int      `json:"offset" db:"-"`
	Limit       *int      `json:"limit" db:"-"`
}
