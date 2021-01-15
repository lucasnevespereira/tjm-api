package api

import (
	"net/http"
	"tjm-api/collector"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to TJM API"))
}

func collectorHandler(w http.ResponseWriter, r *http.Request) {
	collector.Start()
}
