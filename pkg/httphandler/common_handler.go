package httphandler

import "net/http"

// @Summary Hello world.
// @Tags common
// @Success 200 {object} string
// @Router / [get]
func handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}
}
