package utils

import (
	"encoding/json"
	"net/http"
)

// sendJSONResponse is a helper function to send standardized JSON responses.
// It accepts a status code, message, and optional data payload.
func SendJSONResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := map[string]interface{}{
		"message": message,
	}

	if statusCode >= 400 {
		response["success"] = false
	} else {
		response["success"] = true
	}

	if data != nil {
		response["data"] = data
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
