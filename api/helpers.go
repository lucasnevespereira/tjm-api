package api

import (
	"encoding/json"
	"net/http"
)

func retResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	retResponse(w, http.StatusNotFound, map[string]string{"error": "Not found"})
}
