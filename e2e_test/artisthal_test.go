package e2e_test

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

type ArtistHALSuite struct {
	suite.Suite
}

func TestArtistHALSuite(t *testing.T) {
	t.Skip("Not yet")
	suite.Run(t, new(ArtistHALSuite))
}

func (s *ArtistHALSuite) TestEmptyArtists() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db.Exec(
		"INSERT INTO artist (name, image, genres, year_started) VALUES (?, ?, ?, ?)",
		"Slipknot",
		"slipknot.jpg",
		"Nu Metal",
		1995,
	)

	c := http.Client{}

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/artist", nil)
	req.Header.Set("Accept", "application/hal+json")

	resp, err := c.Do(req)

	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal("application/hal+json", resp.Header.Get("Content-Type"))

	expBody := `{"_links": {"self":{"href": "/artist"}}, "_embedded": {"artist": [{"name": "Slipknot", "image": "slipknot.jpg", "genre":"Nu Metal", "started_year": 1995}]}}`
	actBody, _ := ioutil.ReadAll(resp.Body)

	s.JSONEq(expBody, string(actBody))
}
