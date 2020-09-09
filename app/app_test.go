package app_test

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/braddle/go-http-template/app"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/suite"
)

type ApplicationSuite struct {
	suite.Suite
}

func TestApplicationSuite(t *testing.T) {
	suite.Run(t, new(ApplicationSuite))
}

func (s *ApplicationSuite) TestHealthCheck() {
	logBuf := bytes.NewBufferString("")
	log.SetOutput(logBuf)
	log.SetFormatter(&log.JSONFormatter{})

	router := mux.NewRouter()

	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	app.New(router, db)

	url := "/healthcheck"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(`{"status": "OK", "errors": []}`, string(body))

	access := make(map[string]interface{})
	sc := bufio.NewScanner(logBuf)
	sc.Scan()

	json.Unmarshal(sc.Bytes(), &access)

	s.Equal(url, access["request"].(string))
	s.Equal(http.MethodGet, access["method"].(string))
}

func (s *ApplicationSuite) TestNotFound() {
	logBuf := bytes.NewBufferString("")
	log.SetOutput(logBuf)
	log.SetFormatter(&log.JSONFormatter{})

	router := mux.NewRouter()

	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	app.New(router, db)

	url := "/never/going/to/exist"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	s.Equal(http.StatusNotFound, recorder.Code)

	access := make(map[string]interface{})
	sc := bufio.NewScanner(logBuf)
	sc.Scan()

	json.Unmarshal(sc.Bytes(), &access)

	s.Equal(url, access["request"].(string))
	s.Equal(http.StatusNotFound, int(access["status"].(float64)))
}

func (s *ApplicationSuite) TestArtistHAL() {
	// TODO Fix test
	return
	logBuf := bytes.NewBufferString("")
	log.SetOutput(logBuf)
	log.SetFormatter(&log.JSONFormatter{})

	router := mux.NewRouter()

	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	app.New(router, db)

	url := "/artist"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(`{"_links": {"self":{"href": "/artist"}}, "_embedded": {"artist": []}}`, string(body))

	access := make(map[string]interface{})
	sc := bufio.NewScanner(logBuf)
	sc.Scan()

	json.Unmarshal(sc.Bytes(), &access)

	s.Equal(url, access["request"].(string))
	s.Equal(http.MethodGet, access["method"].(string))
}
