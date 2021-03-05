package service

import "net/http"

// @Summary Hello world.
// @Tags common
// @Produce plain
// @Success 200 {object} string
// @Router / [get]
func (s *service) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}
}
