package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dominicbarnes/go-siren"

	"github.com/RichardKnop/jsonhal"
)

type Artists struct {
}

type halArtist struct {
	jsonhal.Hal
}

type artist struct{}

func (h Artists) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var b []byte

	if req.Header.Get("Accept") == "application/siren+json" {
		e := siren.Entity{
			Class: siren.Classes{"artist"},
			Links: []siren.Link{
				{Rel: siren.Rels{"self"}, Href: siren.Href(req.URL.Path)},
			},
			Properties: siren.Properties{
				"size": 0,
			},
			Entities: []siren.EmbeddedEntity{},
		}
		b, _ = json.Marshal(e)
		resp.Header().Set("Content-Type", "application/siren+json")
	} else {
		a := halArtist{}
		a.SetEmbedded("artist", jsonhal.Embedded([]*artist{}))
		a.SetLink("self", req.URL.Path, "")
		b, _ = json.Marshal(a)
		resp.Header().Set("Content-Type", "application/hal+json")
	}

	resp.Write(b)
}