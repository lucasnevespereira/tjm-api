package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Init inits API handlers
func Init() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	http.ListenAndServe(":4000", router)
}
