package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

// RespondJSON escribe data como JSON con el status HTTP indicado.
func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// RespondError escribe un error en formato JSON consistente: {"error": "..."}.
func RespondError(w http.ResponseWriter, status int, mensaje string) {
	RespondJSON(w, status, map[string]string{"error": mensaje})
}

// statusDeError mapea los errores del servicio a códigos HTTP.
func statusDeError(err error) int {
	switch {
	case errors.Is(err, service.ErrCodigoVacio):
		return http.StatusBadRequest
	case errors.Is(err, service.ErrProblemaVacio):
		return http.StatusBadRequest
	case errors.Is(err, service.ErrEtapaVacia):
		return http.StatusBadRequest
	case errors.Is(err, service.ErrDescripcionVacia):
		return http.StatusBadRequest
	case errors.Is(err, service.ErrOrdenIDRequerido):
		return http.StatusBadRequest
	case errors.Is(err, service.ErrNoEncontrado):
		return http.StatusNotFound
	case errors.Is(err, service.ErrTokenInvalido):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
