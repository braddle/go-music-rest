package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ArtistSIRENSuite struct {
	suite.Suite
}

func TestArtistSIRENSuite(t *testing.T) {
	suite.Run(t, new(ArtistSIRENSuite))
}

func (s *ArtistSIRENSuite) TestEmptyArtists() {
	c := http.Client{}

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/artist", nil)
	req.Header.Set("Accept", "application/siren+json")

	resp, err := c.Do(req)

	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal("application/siren+json", resp.Header.Get("Content-Type"))

	expBody := `{"class":["artist"], "links": [{"rel": ["self"], "href": "/artist"}], "properties": {"size": 0}}`
	actBody, _ := ioutil.ReadAll(resp.Body)

	s.JSONEq(expBody, string(actBody))
}
