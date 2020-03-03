package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithError with status code & message
func RespondWithError(w http.ResponseWriter, r *http.Request, code int, message string) {
	RespondWithJSON(w, r, code, map[string]string{
		"msg":   "error",
		"error": message})
}

// RespondWithJSON with status code & data
func RespondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	// LogInfo("respond-1")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
