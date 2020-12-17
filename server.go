package main

import (
	"net/http"

	"github.com/JnecUA/spotify-music-downloader/downloader"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Сделай так, чтобы юникод коды превращались в символы
	r.GET("/get-tracklist-from-playlist-url", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"tracklist": downloader.GetTracklist("https://open.spotify.com/playlist/3DSrtF2DRfwA1wrXlfGy5N")})
	})

	r.Run()
}
