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
	Group       *[]string `form:"group" db:"group"`
	Title       *[]string `form:"song" db:"song"`
	ReleaseDate *[]string `form:"releaseDate" db:"release_date"`
	Before      *string   `form:"before" db:"release_date"`
	After       *string   `form:"after" db:"release_date"`
	Offset      *int      `form:"offset" db:"-"`
	Limit       *int      `form:"limit" db:"-"`
}
