package main

import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"sync"
	"time"
	"math/rand"
	"database/sql"
	"github.com/bariskas/go-course/workshop4/ffmpegUtils"
	"path/filepath"
	_ "github.com/go-sql-driver/mysql"
)

const dirPath = "C:\\Users\\Bariskas\\Desktop\\nginx-1.15.7\\html\\content"

type Task struct {
	id int
}

func GenerateTask(db *sql.DB) *Task {
	var id int
	err := db.QueryRow("SELECT id FROM video WHERE status = 1").Scan(&id)
	if err != nil {
		return nil
	}
	q := `UPDATE video SET status = 2 WHERE id = ?`
	_, err = db.Exec(q, id)
	if err != nil {
		log.Fatal(err)
	}
	return &Task{id}
}

func TaskProvider(db *sql.DB, stopChan chan struct{}) <-chan *Task {
	tasksChan := make(chan *Task)
	go func() {
		for {
			select {
			case <-stopChan:
				close(tasksChan)
				return
			default:
			}
			if task := GenerateTask(db); task != nil {
				log.Printf("got the task %v\n", task.id)
				tasksChan <- task
			} else {
				log.Println("no task for processing, start waiting")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return tasksChan
}

func RunTaskProvider(db * sql.DB, stopChan chan struct{}) <-chan *Task {
	resultChan := make(chan *Task)
	stopTaskProviderChan := make(chan struct{})
	taskProviderChan := TaskProvider(db, stopTaskProviderChan)
	onStop := func () {
		stopTaskProviderChan <- struct{}{}
		close(resultChan)
	}
	go func() {
		for {
			select {
			case <-stopChan:
				onStop()
				return
			case task := <-taskProviderChan:
				select {
				case <-stopChan:
					onStop()
					return
				case resultChan <- task:
				}
			}
		}
	}()
	return resultChan
}

func Worker(db *sql.DB, tasksChan <-chan *Task, name int) {
	log.Printf("start worker %v\n", name)
	for task := range tasksChan {
		log.Printf("start handle task %v on worker %v\n", task.id, name)

		var videoKey string
		var videoName string
		err := db.QueryRow("SELECT video_key, title FROM video WHERE id = ?", task.id).Scan(&videoKey, &videoName)
		if err != nil {
			log.Fatal(err)
		}
		absoluteFilePath := filepath.Join(dirPath, videoKey, videoName)
		absoluteThumbnailPath := filepath.Join(dirPath, videoKey, `screen.jpg`)
		videoDuration, err := ffmpegUtils.GetVideoDuration(absoluteFilePath)
		if err != nil {
			log.Fatal(err)
		}
		err = ffmpegUtils.CreateVideoThumbnail(absoluteFilePath, absoluteThumbnailPath, int64(videoDuration) / 2)
		if err != nil {
			log.Fatal(err)
		}

		q := `UPDATE video SET duration = ?, status = ? WHERE id = ?`
		_, err = db.Exec(q, videoDuration, 3, task.id)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("end handle task %v on worker %v\n", task.id, name)
	}
	log.Printf("stop worker %v\n", name)
}

func RunWorkerPool(db *sql.DB, stopChan chan struct{}) *sync.WaitGroup {
	var wg sync.WaitGroup
	tasksChan := RunTaskProvider(db, stopChan)
	for i := 0; i < 1; i++ {
		go func(i int) {
			wg.Add(1)
			Worker(db, tasksChan, i)
			wg.Done()
		}(i)
	}
	return &wg
}

func main() {
	db, err := sql.Open("mysql", `root:root@tcp(127.0.0.1:3306)/go_workshop_3`)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rand.Seed(time.Now().Unix())
	stopChan := make(chan struct{})
	killChan := getKillSignalChan()
	wg := RunWorkerPool(db, stopChan)

	waitForKillSignal(killChan)
	stopChan <- struct{}{}
	wg.Wait()
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Println("got SIGINT...")
	case syscall.SIGTERM:
		log.Println("got SIGTERM...")
	}
}