package handlers

import (
	"encoding/json"
	"net/http"
)

// RespondJSON escribe una respuesta JSON con el status code dado.
func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// RespondError escribe una respuesta JSON con un mensaje de error.
func RespondError(w http.ResponseWriter, status int, mensaje string) {
	RespondJSON(w, status, map[string]string{"error": mensaje})
}
