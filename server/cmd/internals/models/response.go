package models

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (rep *Response) WriteSuccessResponse(w http.ResponseWriter, message string, statusCode int) {
	rep.Type = "success"
	rep.Message = message
	rep.writeJSON(w, statusCode)
}

func (rep *Response) WriteErrorResponse(w http.ResponseWriter, message string) {
	rep.Type = "error"
	rep.Message = message
	rep.writeJSON(w, http.StatusBadRequest)
}

func (rep *Response) writeJSON(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(rep)

	if err != nil {
		// TODO: log error -- make some app level logger for this
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
