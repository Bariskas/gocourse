package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", listHandler).Methods(http.MethodGet)
	s.HandleFunc("/video/{video_id}", videoHandler).Methods(http.MethodGet)
	return r
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}

func getVideos() []VideoListItem {
	return []VideoListItem{
		VideoListItem{
			"d290f1ee-6c54-4b01-90e6-d701748f0851",
			"Black Retrospetive Woman",
			15,
			"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg"},
		VideoListItem{
			"sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
			"Go Rally TEASER-HD",
			41,
			"/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg"},
		VideoListItem{
			"hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
			"Танцор",
			92,
			"/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg"}}
}
