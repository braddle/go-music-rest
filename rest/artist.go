package rest

import (
	"encoding/json"
	"net/http"

	artist2 "github.com/braddle/go-http-template/artist"

	"github.com/dominicbarnes/go-siren"

	"github.com/RichardKnop/jsonhal"
)

type Artists struct {
	R artist2.ArtistFinder
}

type halArtist struct {
	jsonhal.Hal
}

type artist struct{}

func (h Artists) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var b []byte

	ac := h.R.FindAll(artist2.Filter{})

	if req.Header.Get("Accept") == "application/siren+json" {
		e := siren.Entity{
			Class: siren.Classes{"artist"},
			Links: []siren.Link{
				{Rel: siren.Rels{"self"}, Href: siren.Href(req.URL.Path)},
			},
			Properties: siren.Properties{
				"size": ac.Total(),
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
