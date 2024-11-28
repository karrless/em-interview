package models

type Song struct {
	ID          int64  `json:"id" db:"id"`
	Group       string `json:"group" db:"group"`
	Title       string `json:"song" db:"song"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
	Text        string `json:"text" db:"text"`
	Link        string `json:"link" db:"link"`
}
