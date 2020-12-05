package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
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
	r.HTMLDoc.Find("div[data-testid='tracklist-row']").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find("aria-colindex='2'").Find("span > span").Text())
	})
}
