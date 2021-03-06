package httpresp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DataResponse describes a general JSON response.
type DataResponse struct {
	StatusCode int         `json:"statusCode"`
	Errors     []string    `json:"errors,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func writeMessage(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, `{"statusCode": %d, "message": "%s"}`, statusCode, message)
}

func Error(w http.ResponseWriter, message string, statusCode int) {
	writeMessage(w, message, statusCode)
}

func OK(w http.ResponseWriter) {
	writeMessage(w, "OK", http.StatusOK)
}

// BadRequest replies with HTTP BadRequest code (400).
func BadRequest(w http.ResponseWriter, message string) {
	writeMessage(w, message, http.StatusBadRequest)
}

// NotFound replies with HTTP NOT FOUND code (404).
func NotFound(w http.ResponseWriter, message string) {
	writeMessage(w, message, http.StatusNotFound)
}

// ServerError replies with HTTP InternalServerError code (500).
func ServerError(w http.ResponseWriter, message string) {
	writeMessage(w, message, http.StatusInternalServerError)
}

// Created replies with HTTP CREATED code (201).
func Created(w http.ResponseWriter, message string) {
	writeMessage(w, message, http.StatusCreated)
}

// Accepted replies with HTTP ACCEPTED code (202).
func Accepted(w http.ResponseWriter, message string) {
	writeMessage(w, message, http.StatusAccepted)
}

// NoContent replies with HTTP NO CONTENT code (204).
func NoContent(w http.ResponseWriter) {
	writeMessage(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
}

// Data replies with code and data as a JSON document.
func Data(w http.ResponseWriter, data interface{}, statusCode int) {
	resp := DataResponse{
		StatusCode: statusCode,
		Data:       data,
	}
	body, err := json.Marshal(resp)
	if err != nil {
		Error(w, "Failed to marshal response", http.StatusInternalServerError)
	} else {
		w.WriteHeader(statusCode)
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
