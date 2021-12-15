package models

import "net/http"

type HslServer struct {
	SongsDir string
	Port     int
}

type HslService interface {
	Serve(path string, headers http.HandlerFunc) error
}