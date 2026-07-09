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
// SERVICIOS DIGITALES
// =========================================================

func (s *Server) ListarServicios(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Servicios.Listar())
}

func (s *Server) ObtenerServicio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	servicio, err := s.Servicios.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, "servicio digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, servicio)
}

func (s *Server) CrearServicio(w http.ResponseWriter, r *http.Request) {
	var nuevo models.ServicioDigital
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creado, err := s.Servicios.Crear(nuevo)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarServicio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.ServicioDigital
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	actualizado, err := s.Servicios.Actualizar(id, datos)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			RespondError(w, http.StatusNotFound, "servicio digital no encontrado")
			return
		}
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarServicio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Servicios.Borrar(id); err != nil {
		RespondError(w, http.StatusNotFound, "servicio digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// SUSCRIPCIONES CLIENTES
// =========================================================

func (s *Server) ListarSuscripciones(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Suscripciones.Listar())
}

func (s *Server) ObtenerSuscripcion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	suscripcion, err := s.Suscripciones.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, "suscripción no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, suscripcion)
}

func (s *Server) CrearSuscripcion(w http.ResponseWriter, r *http.Request) {
	var nueva models.SuscripcionCliente
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creada, err := s.Suscripciones.Crear(nueva)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarSuscripcion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.SuscripcionCliente
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	actualizada, err := s.Suscripciones.Actualizar(id, datos)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			RespondError(w, http.StatusNotFound, "suscripción no encontrada")
			return
		}
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) BorrarSuscripcion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Suscripciones.Borrar(id); err != nil {
		RespondError(w, http.StatusNotFound, "suscripción no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// ACCESOS DIGITALES
// =========================================================

func (s *Server) ListarAccesos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Accesos.Listar())
}

func (s *Server) ObtenerAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	acceso, err := s.Accesos.Obtener(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, "acceso digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, acceso)
}

func (s *Server) CrearAcceso(w http.ResponseWriter, r *http.Request) {
	var nuevo models.AccesoDigital
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	creado, err := s.Accesos.Crear(nuevo)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.AccesoDigital
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	actualizado, err := s.Accesos.Actualizar(id, datos)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			RespondError(w, http.StatusNotFound, "acceso digital no encontrado")
			return
		}
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if err := s.Accesos.Borrar(id); err != nil {
		RespondError(w, http.StatusNotFound, "acceso digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
