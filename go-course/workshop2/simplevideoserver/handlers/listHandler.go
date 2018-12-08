package handlers

import (
	"net/http"
	"io"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func listHandler(w http.ResponseWriter, _ *http.Request) {
	videos := getVideos();
	b, err := json.Marshal(videos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
