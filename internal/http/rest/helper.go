package rest

import (
	"encoding/json"
	"net/http"
)

type ResponseJSON struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}

func writeError(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, nil, msg)
}

func writeSuccess(w http.ResponseWriter, code int, body any) {
	writeJSON(w, code, body, "")
}

func writeJSON(w http.ResponseWriter, code int, body any, errMsg string) {
	w.Header().Set("Content-Type", "application/json; chatset=utf-8")
	w.WriteHeader(code)

	response := &ResponseJSON{
		Success: errMsg == "",
		Data:    body,
		Error:   errMsg,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "error encode answer", http.StatusInternalServerError)
	}
}
