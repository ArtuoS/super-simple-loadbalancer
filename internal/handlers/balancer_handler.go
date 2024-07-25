package handlers

import "net/http"

func HandleRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-path", http.StatusMovedPermanently)
	})
}
