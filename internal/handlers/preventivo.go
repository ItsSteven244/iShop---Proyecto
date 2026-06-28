package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

// PreventivoServer guarda el storage que voy a usar en los handlers.
type PreventivoServer struct {
	Storage *storage.MemoriaPreventivo
}

// NewPreventivoServer inicializa el server con el storage recibido.
func NewPreventivoServer(s *storage.MemoriaPreventivo) *PreventivoServer {
	return &PreventivoServer{Storage: s}
}

// PreventivoRouter define todas las rutas de mi módulo preventivo con subrouter Chi.
func PreventivoRouter(store *storage.MemoriaPreventivo) http.Handler {
	s := NewPreventivoServer(store)
	r := chi.NewRouter()

	// Rutas para mantenimientos preventivos
	r.Get("/mantenimientos", s.ListarMantenimientos)
	r.Post("/mantenimientos", s.CrearMantenimiento)
	r.Get("/mantenimientos/{id}", s.ObtenerMantenimiento)
	r.Put("/mantenimientos/{id}", s.ActualizarMantenimiento)
	r.Delete("/mantenimientos/{id}", s.BorrarMantenimiento)

	return r
}

// =========================================================
// MANTENIMIENTOS PREVENTIVOS
// =========================================================

// ListarMantenimientos devuelve todos los mantenimientos registrados.
func (s *PreventivoServer) ListarMantenimientos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarMantenimientos())
}

// ObtenerMantenimiento busca un mantenimiento por su ID.
func (s *PreventivoServer) ObtenerMantenimiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	mantenimiento, encontrado := s.Storage.BuscarMantenimientoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "mantenimiento preventivo no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, mantenimiento)
}

// CrearMantenimiento recibe los datos del mantenimiento, valida los campos obligatorios y lo guarda.
func (s *PreventivoServer) CrearMantenimiento(w http.ResponseWriter, r *http.Request) {
	var nuevo models.MantenimientoPreventivo

	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	if strings.TrimSpace(nuevo.Equipo) == "" {
		RespondError(w, http.StatusBadRequest, "el campo equipo es obligatorio")
		return
	}

	if strings.TrimSpace(nuevo.FechaProgramada) == "" {
		RespondError(w, http.StatusBadRequest, "el campo fecha_programada es obligatorio")
		return
	}

	if strings.TrimSpace(nuevo.TipoMantenimiento) == "" {
		RespondError(w, http.StatusBadRequest, "el campo tipo_mantenimiento es obligatorio")
		return
	}

	RespondJSON(w, http.StatusCreated, s.Storage.CrearMantenimiento(nuevo))
}

// ActualizarMantenimiento reemplaza completamente un mantenimiento existente por su ID.
func (s *PreventivoServer) ActualizarMantenimiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.MantenimientoPreventivo

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	if strings.TrimSpace(datos.Equipo) == "" {
		RespondError(w, http.StatusBadRequest, "el campo equipo es obligatorio")
		return
	}

	actualizado, encontrado := s.Storage.ActualizarMantenimiento(id, datos)

	if !encontrado {
		RespondError(w, http.StatusNotFound, "mantenimiento preventivo no encontrado")
		return
	}

	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarMantenimiento elimina un mantenimiento por su ID..
func (s *PreventivoServer) BorrarMantenimiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if !s.Storage.BorrarMantenimiento(id) {
		RespondError(w, http.StatusNotFound, "mantenimiento preventivo no encontrado")
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
