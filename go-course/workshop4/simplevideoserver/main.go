package main

import (
	"github.com/bariskas/go-course/workshop4/simplevideoserver/handlers"
	"net/http"
	"os"
	log "github.com/sirupsen/logrus"
	"context"
	"os/signal"
	"syscall"
	"database/sql"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}

	db, err := sql.Open("mysql", `root:root@tcp(127.0.0.1:3306)/go_workshop_3`)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	serverUrl := ":8000"
	killSignalChan := getKillSignalChan()
	srv := startServer(serverUrl, db)

	waitForKillSignal(killSignalChan)
	srv.Shutdown(context.Background())

}

func startServer(serverUrl string, db *sql.DB) *http.Server {
	router := handlers.Router(db)
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
