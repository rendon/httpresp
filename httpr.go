package httpr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DataResponse describes a general JSON response.
type DataResponse struct {
	Status int         `json:"status"`
	Errors []string    `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func writeMessage(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"code": %d, "message": "%s"}`, code, message)
}

func Error(w http.ResponseWriter, message string, code int) {
	writeMessage(w, message, code)
}

// BadRequest replies with HTTP BadRequest code (400).
func BadRequest(w http.ResponseWriter, message string) {
	writeMessage(w, message, http.StatusBadRequest)
}

// Created replies with HTTP CREATED code (201).
func Created(w http.ResponseWriter, message string) {
	writeMessage(w, message, http.StatusCreated)
}

// Data replies with code and data as a JSON document.
func Data(w http.ResponseWriter, data interface{}, code int) {
	resp := DataResponse{
		Status: code,
		Data:   data,
	}
	body, err := json.Marshal(resp)
	if err != nil {
		Error(w, "Failed to marshal response", http.StatusInternalServerError)
	} else {
		w.WriteHeader(code)
		fmt.Fprintf(w, "%s", body)
	}
}

func SetStandardAPIHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Length")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Accept-Encoding")
	w.Header().Add("Access-Control-Allow-Headers", "X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Methods", "PUT, DELETE")
}
