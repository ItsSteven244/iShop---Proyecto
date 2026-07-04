package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

type PreventivoServer struct {
	Storage       storage.PreventivoRepository
	Mantenimiento *service.MantenimientoPreventivoService
}

func NewPreventivoServer(s storage.PreventivoRepository, mantenimiento *service.MantenimientoPreventivoService) *PreventivoServer {
	return &PreventivoServer{Storage: s, Mantenimiento: mantenimiento}
}

func PreventivoRouter(store storage.PreventivoRepository, mantenimiento *service.MantenimientoPreventivoService) http.Handler {
	s := NewPreventivoServer(store, mantenimiento)
	r := chi.NewRouter()

	r.Get("/mantenimientos", s.ListarMantenimientos)
	r.Post("/mantenimientos", s.CrearMantenimiento)
	r.Get("/mantenimientos/{id}", s.ObtenerMantenimiento)
	r.Put("/mantenimientos/{id}", s.ActualizarMantenimiento)
	r.Delete("/mantenimientos/{id}", s.BorrarMantenimiento)

	r.Get("/tareas", s.ListarTareas)
	r.Post("/tareas", s.CrearTarea)
	r.Get("/tareas/{id}", s.ObtenerTarea)
	r.Put("/tareas/{id}", s.ActualizarTarea)
	r.Delete("/tareas/{id}", s.BorrarTarea)

	r.Get("/insumos", s.ListarInsumos)
	r.Post("/insumos", s.CrearInsumo)
	r.Get("/insumos/{id}", s.ObtenerInsumo)
	r.Put("/insumos/{id}", s.ActualizarInsumo)
	r.Delete("/insumos/{id}", s.BorrarInsumo)

	return r
}

// =========================================================
// MANTENIMIENTOS (vía service)
// =========================================================

func (s *PreventivoServer) ListarMantenimientos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Mantenimiento.Listar())
}

func (s *PreventivoServer) ObtenerMantenimiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	mant, err := s.Mantenimiento.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, "mantenimiento no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, mant)
}

func (s *PreventivoServer) CrearMantenimiento(w http.ResponseWriter, r *http.Request) {
	var nuevo models.MantenimientoPreventivo
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creado, err := s.Mantenimiento.Crear(nuevo)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

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
	actualizado, err := s.Mantenimiento.Actualizar(id, datos)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			RespondError(w, http.StatusNotFound, "mantenimiento no encontrado")
			return
		}
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *PreventivoServer) BorrarMantenimiento(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Mantenimiento.Borrar(id); err != nil {
		RespondError(w, http.StatusNotFound, "mantenimiento no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// TAREAS (directo al storage)
// =========================================================

func (s *PreventivoServer) ListarTareas(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarTareas())
}

func (s *PreventivoServer) ObtenerTarea(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	tarea, encontrado := s.Storage.BuscarTareaPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "tarea no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, tarea)
}

func (s *PreventivoServer) CrearTarea(w http.ResponseWriter, r *http.Request) {
	var nueva models.TareaPreventiva
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nueva.Descripcion) == "" {
		RespondError(w, http.StatusBadRequest, "el campo descripcion es obligatorio")
		return
	}
	if nueva.MantenimientoPreventivoID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo mantenimiento_preventivo_id es obligatorio")
		return
	}
	RespondJSON(w, http.StatusCreated, s.Storage.CrearTarea(nueva))
}

func (s *PreventivoServer) ActualizarTarea(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.TareaPreventiva
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	actualizada, encontrado := s.Storage.ActualizarTarea(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "tarea no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *PreventivoServer) BorrarTarea(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if !s.Storage.BorrarTarea(id) {
		RespondError(w, http.StatusNotFound, "tarea no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// INSUMOS (directo al storage)
// =========================================================

func (s *PreventivoServer) ListarInsumos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarInsumos())
}

func (s *PreventivoServer) ObtenerInsumo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	insumo, encontrado := s.Storage.BuscarInsumoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "insumo no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, insumo)
}

func (s *PreventivoServer) CrearInsumo(w http.ResponseWriter, r *http.Request) {
	var nuevo models.InsumoPreventivo
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if nuevo.MantenimientoPreventivoID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo mantenimiento_preventivo_id es obligatorio")
		return
	}
	RespondJSON(w, http.StatusCreated, s.Storage.CrearInsumo(nuevo))
}

func (s *PreventivoServer) ActualizarInsumo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.InsumoPreventivo
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	actualizado, encontrado := s.Storage.ActualizarInsumo(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "insumo no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *PreventivoServer) BorrarInsumo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if !s.Storage.BorrarInsumo(id) {
		RespondError(w, http.StatusNotFound, "insumo no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
