package spotify_tracklist

import (
	"fmt"
	"os"
	"strings"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func GetTracklist(url string) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(url, g.Opt.ParseFunc)
		},
		ParseFunc: parseSongs,
	}).Start()
}

func parseSongs(g *geziyor.Geziyor, r *client.Response) {
	track_list := strings.Split(strings.Split(string(r.Body), `"tracks"`)[1], ";") //Get track list in json format
	track_list = strings.Split(strings.TrimSpace(track_list[0]), `"track":{`)      //Split tracks
	track_list = track_list[1:len(track_list)]
	track_list = jsonToSongName(track_list)
	fmt.Println(track_list)
}

func jsonToSongName(track_list []string) []string {
	var new_track_list []string
	for _, track := range track_list {
		new_track_list = append(new_track_list, findSongName(track)+" - "+findArtistsNames(track))
	}
	return new_track_list
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

func ParseVideo(g *geziyor.Geziyor, r *client.Response) {
	f, _ := os.Create("lol.txt")
	f.Write(r.Body)
}
