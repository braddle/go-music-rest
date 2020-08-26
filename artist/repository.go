package artist

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db}
}

type Filter struct{}

func (r Repository) FindAll(f Filter) ArtistCollection {
	return ArtistCollection{}
}

type ArtistCollection struct{}
type Artist struct{}

func (c ArtistCollection) Total() int {
	return 0
}

func (c ArtistCollection) Collection() []Artist {
	return []Artist{}
}
