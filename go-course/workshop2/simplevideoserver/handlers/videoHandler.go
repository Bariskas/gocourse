package handlers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func videoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	video, _ := getVideoById(vars["video_id"])
	jsonOutput := `{"id": "%s",
		"name": "%s",
		"duration": %d,
		"thumbnail":"/content/%s/screen.jpg",
		"url":"/content/%s/index.mp4"}`
	fmt.Fprintf(w, jsonOutput, vars["video_id"], video.Name, video.Duration, vars["video_id"], vars["video_id"])
}

func getVideoById(id string) (*VideoListItem, int) {
	for _, video := range getVideos() {
		if video.Id == id {
			return &video, 0
		}
	}
	return nil, 1
}
