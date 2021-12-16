package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	youtube "github.com/kkdai/youtube/v2"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func dl() {
	//youtube.ExtractVideoID("https://www.youtube.com/watch?v=rFejpH_tAHM")
	videoID := "BaW_jenozKc"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := create("./songs/video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

}

func Convert() {
	cmd := &exec.Cmd{
		Path:   "./convert.sh",
		Args:   []string{"./convert"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	cmd.Start()
	cmd.Wait()

}

func main() {
	// une vague tentative de faire de la concurrence...
	// faut que j'approfondisse cela
	go dl()
	Convert()
	// configure the songs directory name and port
	const songsDir = "songs"
	const port = 8080

	// add a handler for the song files
	http.Handle("/", addHeaders(http.FileServer(http.Dir(songsDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", songsDir, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
