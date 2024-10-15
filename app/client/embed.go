package client

import (
	"embed"
	"net/http"
)

//go:embed index.html
var indexHtmlFile []byte

func EmbedIndexHtml() []byte {
	return indexHtmlFile
}

//go:embed main.js
var JavaScriptFiles embed.FS

func EmbedFileServerHandler(prefix string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.FS(JavaScriptFiles)))
}
