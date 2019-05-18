package model

import (
	"encoding/json"
	"net/http"
)

// Writes the response as a standard JSON response with StatusOK
func WriteOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&jsonResponse{Success: true, Status: http.StatusOK, Data: m}); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Sever Error!")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func WriteErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).Encode(&jsonErrorResponse{Success: false, Status: errorCode, Error: errorMsg})
}
