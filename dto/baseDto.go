package dto

import (
	"encoding/json"
	"net/http"
)

func ToJsonResponse(w http.ResponseWriter, data interface{}, status int) {
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

func ToStatusResponse(w http.ResponseWriter, status int) {
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func FromRequest(r *http.Request, data interface{}) {
	_ = json.NewDecoder(r.Body).Decode(data)
}
