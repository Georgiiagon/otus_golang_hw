package internalhttp

import (
	"net/http"
)

func (s *Server) loggingMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("hello there!")
		h(w, r)
	}
}
