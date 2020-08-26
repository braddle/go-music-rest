package artist_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/braddle/go-http-template/artist"

	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "tester"
	password = "testing"
	dbname   = "music"
)

func (s *ArtistRepositorySuite) TestNoneAvailableForFindAll() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, _ := sql.Open("postgres", psqlInfo)

	ar := artist.NewRepository(db)

	ac := ar.FindAll(artist.Filter{})

	s.Zero(ac.Total())
	s.Empty(ac.Collection())
}

//func (s *ArtistRepositorySuite) TestUnfilteredFindAll() {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//
//	db, _ := sql.Open("postgres", psqlInfo)
//
//	sql := `INSERT INTO artist
//				()
//			VALUES
//				(),`
//
//	db.Exec(sql)
//
//	ar := artist.NewRepository(db)
//
//	ac := ar.FindAll(artist.Filter{})
//
//	s.Equal(2, ac.Total())
//
//	artists := []artist.Artist{
//		{},
//		{},
//	}
//	s.Equal(artists, ac.Collection())
//}

type ArtistRepositorySuite struct {
	suite.Suite
}

func TestArtistRepositorySuite(t *testing.T) {
	suite.Run(t, new(ArtistRepositorySuite))
}
