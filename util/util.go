package util

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func JsonResponse(w http.ResponseWriter, data interface{}, status int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Write the JSON response to the response writer
	w.Write(jsonData)
}
