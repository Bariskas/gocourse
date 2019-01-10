package handlers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getVideos(db *sql.DB) []VideoListItem {
	var videos []VideoListItem

	rows, err := db.Query("SELECT video_key, title, duration, thumbnail_url, url FROM video")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var video VideoListItem
		err := rows.Scan(&video.Id, &video.Name, &video.Duration, &video.Thumbnail, &video.Url)
		if err != nil {
			return videos
		}
		videos = append(videos, video)
	}
	return videos
}

func insertFileInfo(db *sql.DB, fileName string, uuidString string) {
	q := `INSERT INTO video (video_key, title, status, duration, url, thumbnail_url) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(q, uuidString, fileName, 3, 100, `/content/` + uuidString + `/` + fileName, `/content/` + uuidString + `/screen.jpg`)
	if err != nil {
		log.Fatal(err)
	}
}