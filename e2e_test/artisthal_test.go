package e2e

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ArtistHALSuite struct {
	suite.Suite
}

func TestArtistHALSuite(t *testing.T) {
	suite.Run(t, new(ArtistHALSuite))
}

func (s *ArtistHALSuite) TestEmptyArtists() {
	c := http.Client{}

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/artist", nil)
	req.Header.Set("Accept", "application/hal+json")

	resp, err := c.Do(req)

	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal("application/hal+json", resp.Header.Get("Content-Type"))

	expBody := `{"_links": {"self":{"href": "/artist"}}, "_embedded": {"artist": []}}`
	actBody, _ := ioutil.ReadAll(resp.Body)

	s.JSONEq(expBody, string(actBody))
}
