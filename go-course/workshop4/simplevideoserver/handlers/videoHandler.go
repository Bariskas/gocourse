package handlers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"database/sql"
)

func videoHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	video, _ := getVideoById(db, vars["video_id"])
	jsonOutput := `{"id": "%s",
		"name": "%s",
		"duration": %d,
		"thumbnail":"/content/%s/screen.jpg",
		"url":"/content/%s/%s"}`
	fmt.Fprintf(w, jsonOutput, vars["video_id"], video.Name, video.Duration, vars["video_id"], vars["video_id"],  video.Name)
}

func getVideoById(db *sql.DB, id string) (*VideoListItem, int) {
	for _, video := range getVideos(db) {
		if video.Id == id {
			return &video, 0
		}
	}
	return nil, 1
}
