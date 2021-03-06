package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	log "github.com/sirupsen/logrus"
	"database/sql"
)

func Router(db *sql.DB) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", wrapHandlerWithDB(db, listHandler)).Methods(http.MethodGet)
	s.HandleFunc("/video/{video_id}", wrapHandlerWithDB(db, videoHandler)).Methods(http.MethodGet)
	s.HandleFunc("/video", wrapHandlerWithDB(db, uploadVideoHandler)).Methods(http.MethodPost)
	return logMiddleware(r)
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

func wrapHandlerWithDB(db *sql.DB, f func(*sql.DB, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
}
