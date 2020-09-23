package artist

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

type ArtistFinder interface {
	FindAll(f Filter) ArtistCollection
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db}
}

type Filter struct{}

func (r Repository) FindAll(f Filter) ArtistCollection {
	rows, _ := r.db.Query("SELECT name, genres, image, year_started FROM artist")
	ac := ArtistCollection{}

	for rows.Next() {
		a := Artist{}
		rows.Scan(&a.Name, &a.Genre, &a.Image, &a.Started)

		ac.a = append(ac.a, a)
	}

	return ac
}

type ArtistCollection struct {
	a []Artist
}
type Artist struct {
	Name    string
	Image   string
	Genre   string
	Started int
}

func (c ArtistCollection) Total() int {
	return len(c.a)
}

func (c ArtistCollection) Collection() []Artist {
	return c.a
}
