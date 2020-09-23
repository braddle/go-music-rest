package e2e_test

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

var db *sql.DB
var c http.Client

func (s *ArtistSIRENSuite) SetupTest() {
	db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db.Exec("TRUNCATE artist")

	c = http.Client{}

}
func (s *ArtistSIRENSuite) TestEmptyArtists() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/artist", nil)
	req.Header.Set("Accept", "application/siren+json")

	resp, err := c.Do(req)

	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal("application/siren+json", resp.Header.Get("Content-Type"))

	expBody := `{
	"class":["artist"], 
	"links": [{"rel": ["self"], "href": "/artist"}], 
	"properties": {"size": 0}
}`
	actBody, _ := ioutil.ReadAll(resp.Body)

	s.JSONEq(expBody, string(actBody))
}

func (s *ArtistSIRENSuite) TestManyArtists() {
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

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/artist", nil)
	req.Header.Set("Accept", "application/siren+json")
	resp, err := c.Do(req)

	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal("application/siren+json", resp.Header.Get("Content-Type"))

	expBody := `{
	"class":["artist"], 
	"links": [{"rel": ["self"], "href": "/artist"}], 
	"properties": {"size": 2},
	"entities": [
		{
			"properties" : {
				"name": "Slipknot",
				"image": "slipknot.jpg",
				"genre": "Nu Metal",
				"started": 1995
			},
			"rel": null
		},
		{
			"properties" : {
				"name": "Limp Bizkit",
				"image": "limp-bizkit.jpg",
				"genre": "Nu Metal",
				"started": 1994
			},
			"rel": null
		}
	]
}`
	actBody, _ := ioutil.ReadAll(resp.Body)

	s.JSONEq(expBody, string(actBody))
}

type ArtistSIRENSuite struct {
	suite.Suite
}

func TestArtistSIRENSuite(t *testing.T) {
	suite.Run(t, new(ArtistSIRENSuite))
}
