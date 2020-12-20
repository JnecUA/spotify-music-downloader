package downloader

import (
	"strings"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

var tracklistglobal string

//GetTracklist list of tracks and artists
func GetTracklist(url string) string {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(url, g.Opt.ParseFunc)
		},
		ParseFunc: parseSongs,
	}).Start()
	return tracklistglobal

}

func parseSongs(g *geziyor.Geziyor, r *client.Response) {
	tracklist := strings.Split(strings.Split(string(r.Body), `Spotify.Entity = `)[1], ";")[0] //Get track list in json format
	tracklistglobal = tracklist
}
