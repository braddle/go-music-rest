package rest

import (
	"encoding/json"
	"net/http"

	"github.com/braddle/go-http-template/artist"

	"github.com/dominicbarnes/go-siren"

	"github.com/RichardKnop/jsonhal"
)

type Artists struct {
	R artist.ArtistFinder
}

type halArtist struct {
	jsonhal.Hal
}

func (h Artists) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var b []byte

	ac := h.R.FindAll(artist.Filter{})

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

		for _, a := range ac.Collection() {
			ee := siren.EmbeddedEntity{
				Entity: siren.Entity{
					//Entities:   nil,
					//Links:      nil,
					//Actions:    nil,
					Properties: siren.Properties{
						"name":    a.Name,
						"image":   a.Image,
						"genre":   a.Genre,
						"started": a.Started,
					},
					//Title:      "",
					//Class:      nil,
				},
				//Rel:    nil,
				//Href:   "",
			}
			e.Entities = append(e.Entities, ee)
		}
		b, _ = json.Marshal(e)
		resp.Header().Set("Content-Type", "application/siren+json")
	} else {
		a := halArtist{}
		a.SetEmbedded("artist", jsonhal.Embedded([]*artist.Artist{}))
		a.SetLink("self", req.URL.Path, "")
		b, _ = json.Marshal(a)
		resp.Header().Set("Content-Type", "application/hal+json")
	}

	resp.Write(b)
}
