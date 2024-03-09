package dto

import "net/http"

type MessageResponse struct {
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

func RespondWith400(w http.ResponseWriter) {
	ToJsonResponse(w, MessageResponse{Message: http.StatusText(http.StatusBadRequest)}, http.StatusBadRequest)
}

func RespondWith401(w http.ResponseWriter) {
	ToJsonResponse(w, MessageResponse{Message: http.StatusText(http.StatusUnauthorized)}, http.StatusUnauthorized)
}

func RespondWith403(w http.ResponseWriter) {
	ToJsonResponse(w, MessageResponse{Message: http.StatusText(http.StatusForbidden)}, http.StatusForbidden)
}
