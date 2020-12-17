from __future__ import unicode_literals
import youtube_dl
import urllib.request
import urllib.parse
import re
import requests
import json
import subprocess

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

def GetTracklist(url):
    process = subprocess.Popen("go run server.go")
    try:
        response = requests.get(url)
 
        # если ответ успешен, исключения задействованы не будут
        response.raise_for_status()
    except HTTPError as http_err:
        print(f'HTTP error occurred: {http_err}')  # Python 3.6
    except Exception as err:
        print(f'Other error occurred: {err}')  # Python 3.6
    else:
        return json.loads(response.content)["tracklist"]
    return []

if __name__ == "__main__":
    print("Strating")
    tracklist = GetTracklist('http://localhost:8080/get-tracklist-from-playlist-url')
    for track in tracklist:
        downloadVideo(findVideoLink(track))