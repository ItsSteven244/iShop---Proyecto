package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

// =========================================================
// MANTENIMIENTOS.
// =========================================================

func (s *Server) ListarMantenimientos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Mantenimiento.Listar())
}

func (s *Server) ObtenerMantenimiento(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) CrearMantenimiento(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) ActualizarMantenimiento(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) BorrarMantenimiento(w http.ResponseWriter, r *http.Request) {
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
// TAREAS
// =========================================================

func (s *Server) ListarTareas(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Tareas.Listar())
}

func (s *Server) ObtenerTarea(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	tarea, err := s.Tareas.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, "tarea no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, tarea)
}

func (s *Server) CrearTarea(w http.ResponseWriter, r *http.Request) {
	var nueva models.TareaPreventiva
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creada, err := s.Tareas.Crear(nueva)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarTarea(w http.ResponseWriter, r *http.Request) {
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
	actualizada, err := s.Tareas.Actualizar(id, datos)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			RespondError(w, http.StatusNotFound, "tarea no encontrada")
			return
		}
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) BorrarTarea(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Tareas.Borrar(id); err != nil {
		RespondError(w, http.StatusNotFound, "tarea no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// INSUMOS
// =========================================================

func (s *Server) ListarInsumos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Insumos.Listar())
}

func (s *Server) ObtenerInsumo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	insumo, err := s.Insumos.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, "insumo no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, insumo)
}

func (s *Server) CrearInsumo(w http.ResponseWriter, r *http.Request) {
	var nuevo models.InsumoPreventivo
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creado, err := s.Insumos.Crear(nuevo)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarInsumo(w http.ResponseWriter, r *http.Request) {
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
	actualizado, err := s.Insumos.Actualizar(id, datos)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			RespondError(w, http.StatusNotFound, "insumo no encontrado")
			return
		}
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarInsumo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Insumos.Borrar(id); err != nil {
		RespondError(w, http.StatusNotFound, "insumo no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
