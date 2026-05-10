package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Res struct {
	Message  string  `json:"message"`
	Data 	 any 	 `json:"data"`
}

type Err struct {
	Error string `json:"error"`
}

// Error writes a structured JSON success response
func JSON(w http.ResponseWriter, message string, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := Res{
		Message: message,
		Data: data,
	}

	json.NewEncoder(w).Encode(res)
}


// Error writes a structured JSON error response
func Error(w http.ResponseWriter, message string, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err != nil {
        slog.Error(message, "status", status, "error", err)
    }
	
	errRes := Err{Error: message}

	json.NewEncoder(w).Encode(errRes)
}