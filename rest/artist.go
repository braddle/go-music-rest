package rest

import (
	"encoding/json"
	"net/http"

	"github.com/RichardKnop/jsonhal"
)

type Artists struct {
}

type halArtist struct {
	jsonhal.Hal
}

type artist struct{}

func (h Artists) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	a := halArtist{}
	a.SetEmbedded("artist", jsonhal.Embedded([]*artist{}))
	a.SetLink("self", req.URL.Path, "")

	b, _ := json.Marshal(a)

	resp.Header().Set("Content-Type", "application/hal+json")
	resp.Write(b)
}
