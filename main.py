from __future__ import unicode_literals
import youtube_dl
import urllib.request
import urllib.parse
import re


def findVideoLink(search_query):
    html = urllib.request.urlopen("https://www.youtube.com/results?search_query=" + urllib.parse.quote_plus(search_query, safe=''))
    video_ids = re.findall(r"watch\?v=(\S{11})", html.read().decode())
    return "https://www.youtube.com/watch?v=" + video_ids[0]


def downloadVideo(url):
    ydl_opts = {
        'format': 'bestaudio/best',
    }
    with youtube_dl.YoutubeDL(ydl_opts) as ydl:
        ydl.download([url])

if __name__ == "__main__":
    print("Strating")