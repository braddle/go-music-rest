package rest_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/braddle/go-http-template/rest"
	"github.com/stretchr/testify/suite"
)

type ArtistSuite struct {
	suite.Suite
}

func TestArtistsSuite(t *testing.T) {
	suite.Run(t, new(ArtistSuite))
}

func (s *ArtistSuite) TestNoArtistsHAL() {
	h := rest.Artists{}

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/artist", nil)
	h.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusOK, recorder.Code)
	s.Equal("application/hal+json", recorder.Header().Get("Content-Type"))
	s.JSONEq(`{"_links": {"self":{"href": "/artist"}}, "_embedded": {"artist": []}}`, string(body))
}
