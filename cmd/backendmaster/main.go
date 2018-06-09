package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/undeadops/utils/lib"
)

var (
	logger *lib.Logger
)

func init() {
	logger = lib.NewLogger()
	prometheus.MustRegister(writeEvent)
}

func main() {
	logger.LogInfo("Started in environment: ", os.Getenv("ENV"), "\n")

	cancelChan := lib.MakeCancelChan()

	startupWrite()

	go writeFile(cancelChan)

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/metrics", promhttp.Handler())

	go func() {
		logger.LogInfo("Starting Prometheus Metrics Endpoint")
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			logger.LogFatal("Unable to Start Prometheus Endpoint")
		}
	}()

	for {
		select {
		case <-cancelChan:
			logger.LogInfo("Cancel channel received")
			time.Sleep(5 * time.Second)
			logger.LogInfo("Slept time", "\n")
			finalWrite()
			return
		}
	}
}

func writeFile(cancelChan chan struct{}) {
	file, err := os.Create("results.txt")
	if err != nil {
		logger.LogFatal("Cannot Create File", "\n")
	}
	defer file.Close()

	tick := time.Tick(1500 * time.Millisecond)
	for {
		select {
		case <-tick:
			t := time.Now()
			fmt.Fprintf(file, "%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			//fmt.Fprintf(file, "Hello Readers of golangcode.com")
			writeEvent.With(prometheus.Labels{"filename": "results.txt"}).Inc()
		case <-cancelChan:
			logger.LogInfo("Received Cancel in WriteFile")
			return
		}
	}
}

func startupWrite() {
	file, err := os.Create("results.txt")
	if err != nil {
		logger.LogFatal("Cannot Open File to write", "\n")
	}
	defer file.Close()

	fmt.Fprintf(file, "====== OPENING FILE =======\n")
}

func finalWrite() {
	file, err := os.Create("results.txt")
	if err != nil {
		logger.LogFatal("Cannot Open File to write", "\n")
	}
	defer file.Close()

	fmt.Fprintf(file, "====== CLOSING FILE =======\n")
}
