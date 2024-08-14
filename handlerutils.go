package main

import "net/http"

func RespondWithText(w http.ResponseWriter, statusCode int, text string) {
	respond(w, "text/plain; charset=utf-8", statusCode, text)
}

func respond(w http.ResponseWriter, contentType string, statusCode int, content string) {
	respondBytes(w, contentType, statusCode, []byte(content))
}

func respondBytes(w http.ResponseWriter, contentType string, statusCode int, content []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(content)
	if err != nil {
		http.Error(w, "Internal server error", 500)
	}
}
