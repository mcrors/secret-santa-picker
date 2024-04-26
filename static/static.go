package static

import (
	"embed"
	"io/fs"
)

//go:embed css/* images/*
var content embed.FS

func Content() fs.FS {
	return content
}
