package main

import (
	"io"
	"net/http"
)

func health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	_, _ = io.WriteString(w, `{"status":"ok"}`)
}
