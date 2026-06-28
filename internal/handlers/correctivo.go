package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
)

// =========================================================
// ORDENES CORRECTIVAS
// =========================================================

func (s *Server) ListarOrdenes(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Ordenes.Listar())
}

func (s *Server) ObtenerOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	orden, err := s.Ordenes.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, orden)
}

func (s *Server) CrearOrden(w http.ResponseWriter, r *http.Request) {
	var nueva models.OrdenCorrectiva
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creada, err := s.Ordenes.Crear(nueva)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.OrdenCorrectiva
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	actualizada, err := s.Ordenes.Actualizar(id, datos)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) ActualizarEstadoOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var body struct {
		Estado      string `json:"estado"`
		Diagnostico string `json:"diagnostico"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	actualizada, err := s.Ordenes.ActualizarParcial(id, body.Estado, body.Diagnostico)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) BorrarOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Ordenes.Borrar(id); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// PROCESOS REPARACION
// =========================================================

func (s *Server) ListarProcesos(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Procesos.Listar())
}

func (s *Server) ObtenerProceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	proceso, err := s.Procesos.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, proceso)
}

func (s *Server) CrearProceso(w http.ResponseWriter, r *http.Request) {
	var nuevo models.ProcesoReparacion
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creado, err := s.Procesos.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarProceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.ProcesoReparacion
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	actualizado, err := s.Procesos.Actualizar(id, datos)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarProceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Procesos.Borrar(id); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// EVIDENCIAS DANIO
// =========================================================

func (s *Server) ListarEvidencias(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Evidencias.Listar())
}

func (s *Server) ObtenerEvidencia(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	evidencia, err := s.Evidencias.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, evidencia)
}

func (s *Server) CrearEvidencia(w http.ResponseWriter, r *http.Request) {
	var nueva models.EvidenciaDanio
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creada, err := s.Evidencias.Crear(nueva)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarEvidencia(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.EvidenciaDanio
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	actualizada, err := s.Evidencias.Actualizar(id, datos)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) BorrarEvidencia(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Evidencias.Borrar(id); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}