package downloader

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

var tracklistglobal []string

//GetTracklist list of tracks and artists
func GetTracklist(url string) []string {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(url, g.Opt.ParseFunc)
		},
		ParseFunc: parseSongs,
	}).Start()
	return tracklistglobal

}

func parseSongs(g *geziyor.Geziyor, r *client.Response) {
	tracklist := strings.Split(strings.Split(string(r.Body), `"tracks"`)[1], ";") //Get track list in json format
	tracklist = strings.Split(strings.TrimSpace(tracklist[0]), `"track":{`)       //Split tracks
	tracklist = tracklist[1:]
	tracklist = jsonToSongName(tracklist)
	tracklistglobal = tracklist
}

func jsonToSongName(tracklist []string) []string {
	var newTrackList []string
	for _, track := range tracklist {
		newTrack, err := strconv.Unquote(`"` + findSongName(track) + " - " + findArtistsNames(track) + `"`)
		if err != nil {
			fmt.Println(err.Error())
		}
		newTrackList = append(newTrackList, newTrack)
	}
	return newTrackList
}

// findArtistsNames find artists names from first song in list
func findArtistsNames(l string) string {
	artists := strings.Split(l, `"artists"`)[1]
	artistsNames := strings.Split(artists, "name")
	names := ""
	for i := 1; i < len(artistsNames)-1; i++ {
		names += strings.Split(artistsNames[i], `"`)[2]
		if i != len(artistsNames)-2 {
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
