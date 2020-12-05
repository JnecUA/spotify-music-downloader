package main

import (
	"strings"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://open.spotify.com/playlist/2lfAwrTGw2rbXLJf7elpo9", g.Opt.ParseFunc)
		},
		ParseFunc: parseSongs,
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
}

func parseSongs(g *geziyor.Geziyor, r *client.Response) {
	tracks_list := strings.Split(strings.Split(string(r.Body), `"tracks"`)[1], ";") //Get track list in json format
	tracks_list = strings.Split(strings.TrimSpace(tracks_list[0]), `"track":{`)     //Split tracks
	tracks_list = tracks_list[1:len(tracks_list)]                                   //Delete first elemet without track

}

// findArtistsNames find artists names from first song in list
func findArtistsNames(l string) string {
	artists := strings.Split(l, `"artists"`)[1]
	artists_names := strings.Split(artists, "name")
	names := ""
	for i := 1; i < len(artists_names)-1; i++ {
		names += strings.Split(artists_names[i], `"`)[2]
		if i != len(artists_names)-2 {
			names += ", "
		}
	}
	return names
}

//findSongName find name of first song in list
func findSongName(l string) string {
	track := strings.Split(l, `"artists"`)[1]
	song := strings.Split(track, "name")
	song = strings.Split(song[len(song)-1], `"`)
	return song[2]
}
