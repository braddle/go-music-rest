package artist_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/braddle/go-http-template/artist"

	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

var db *sql.DB

func (s *ArtistRepositorySuite) SetupTest() {
	db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db.Exec("TRUNCATE artist")
}

func (s *ArtistRepositorySuite) TestNoneAvailableForFindAll() {
	ar := artist.NewRepository(db)

	ac := ar.FindAll(artist.Filter{})

	s.Zero(ac.Total())
	s.Empty(ac.Collection())
}

func (s *ArtistRepositorySuite) TestUnfilteredFindAll() {
	name1 := "Slipknot"
	image1 := "slipknot.jpg"
	genre1 := "Nu Metal"
	year1 := 1995

	name2 := "Limp Bizkit"
	image2 := "limp-bizkit.jpg"
	genre2 := "Nu Metal"
	year2 := 1994

	sql := "INSERT INTO artist (name, image, genres, year_started) VALUES ($1, $2, $3, $4)"

	db.Exec(sql, name1, image1, genre1, year1)
	db.Exec(sql, name2, image2, genre2, year2)

	ar := artist.NewRepository(db)

	ac := ar.FindAll(artist.Filter{})

	s.Equal(2, ac.Total())

	artists := []artist.Artist{
		{name1, image1, genre1, year1},
		{name2, image2, genre2, year2},
	}
	s.Equal(artists, ac.Collection())
}

type ArtistRepositorySuite struct {
	suite.Suite
}

func TestArtistRepositorySuite(t *testing.T) {
	suite.Run(t, new(ArtistRepositorySuite))
}
