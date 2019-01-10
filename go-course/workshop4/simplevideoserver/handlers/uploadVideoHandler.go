package handlers

import (
	"net/http"
	"io"
	"os"
	"path/filepath"
	"github.com/google/uuid"
	"database/sql"
	"log"
)

const dirPath = "C:\\Users\\Bariskas\\Desktop\\nginx-1.15.7\\html\\content"

func uploadVideoHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fileReader, header, err := r.FormFile("file[]")
	fileName := header.Filename
	uuidString := uuid.New().String()
	file, err := createFile(fileName, uuidString)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(file, fileReader)
	if err != nil {
		log.Fatal(err)
	}

	absoluteFilePath := filepath.Join(dirPath, uuidString, fileName)
	absoluteThumbnailPath := filepath.Join(dirPath, uuidString, `screen.jpg`)
	insertNewFileInfo(db, fileName, uuidString, absoluteFilePath, absoluteThumbnailPath)
}

func createFile(fileName string, uuidString string) (*os.File, error) {
	if err := os.MkdirAll(dirPath + string(os.PathSeparator) + uuidString, os.ModePerm); err != nil {
		return nil, err
	}
	filePath := filepath.Join(dirPath, uuidString, fileName)
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}