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

type SuscripcionesServer struct {
	Storage   storage.SuscripcionesRepository
	Servicios *service.ServicioDigitalService
}

func NewSuscripcionesServer(s storage.SuscripcionesRepository, servicios *service.ServicioDigitalService) *SuscripcionesServer {
	return &SuscripcionesServer{Storage: s, Servicios: servicios}
}

// SuscripcionesRoutes registra las rutas del módulo Suscripciones directamente
// sobre el router recibido (r), en vez de crear un sub-router propio.
// Esto evita el conflicto de chi al montar dos handlers en el mismo path "/".
func SuscripcionesRoutes(r chi.Router, store storage.SuscripcionesRepository, servicios *service.ServicioDigitalService) {
	s := NewSuscripcionesServer(store, servicios)

	r.Get("/servicios", s.ListarServicios)
	r.Post("/servicios", s.CrearServicio)
	r.Get("/servicios/{id}", s.ObtenerServicio)
	r.Put("/servicios/{id}", s.ActualizarServicio)
	r.Delete("/servicios/{id}", s.BorrarServicio)

	r.Get("/suscripciones", s.ListarSuscripciones)
	r.Post("/suscripciones", s.CrearSuscripcion)
	r.Get("/suscripciones/{id}", s.ObtenerSuscripcion)
	r.Put("/suscripciones/{id}", s.ActualizarSuscripcion)
	r.Delete("/suscripciones/{id}", s.BorrarSuscripcion)

	r.Get("/accesos", s.ListarAccesos)
	r.Post("/accesos", s.CrearAcceso)
	r.Get("/accesos/{id}", s.ObtenerAcceso)
	r.Put("/accesos/{id}", s.ActualizarAcceso)
	r.Delete("/accesos/{id}", s.BorrarAcceso)
}

func (s *SuscripcionesServer) ListarServicios(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Servicios.Listar())
}

func (s *SuscripcionesServer) ObtenerServicio(w http.ResponseWriter, r *http.Request) {
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

func (s *SuscripcionesServer) CrearServicio(w http.ResponseWriter, r *http.Request) {
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

func (s *SuscripcionesServer) ActualizarServicio(w http.ResponseWriter, r *http.Request) {
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

func (s *SuscripcionesServer) BorrarServicio(w http.ResponseWriter, r *http.Request) {
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

func (s *SuscripcionesServer) ListarSuscripciones(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarSuscripciones())
}

func (s *SuscripcionesServer) ObtenerSuscripcion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	suscripcion, encontrado := s.Storage.BuscarSuscripcionPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "suscripción no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, suscripcion)
}

func (s *SuscripcionesServer) CrearSuscripcion(w http.ResponseWriter, r *http.Request) {
	var nueva models.SuscripcionCliente
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if nueva.ClienteID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo cliente_id es obligatorio")
		return
	}
	if nueva.ServicioDigitalID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo servicio_digital_id es obligatorio")
		return
	}
	if strings.TrimSpace(nueva.FechaInicio) == "" {
		RespondError(w, http.StatusBadRequest, "el campo fecha_inicio es obligatorio")
		return
	}
	if strings.TrimSpace(nueva.FechaFin) == "" {
		RespondError(w, http.StatusBadRequest, "el campo fecha_fin es obligatorio")
		return
	}
	if strings.TrimSpace(nueva.Estado) == "" {
		RespondError(w, http.StatusBadRequest, "el campo estado es obligatorio")
		return
	}
	if nueva.TecnicoID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo tecnico_id es obligatorio")
		return
	}
	RespondJSON(w, http.StatusCreated, s.Storage.CrearSuscripcion(nueva))
}

func (s *SuscripcionesServer) ActualizarSuscripcion(w http.ResponseWriter, r *http.Request) {
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
	actualizada, encontrado := s.Storage.ActualizarSuscripcion(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "suscripción no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *SuscripcionesServer) BorrarSuscripcion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if !s.Storage.BorrarSuscripcion(id) {
		RespondError(w, http.StatusNotFound, "suscripción no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

func (s *SuscripcionesServer) ListarAccesos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarAccesos())
}

func (s *SuscripcionesServer) ObtenerAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	acceso, encontrado := s.Storage.BuscarAccesoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "acceso digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, acceso)
}

func (s *SuscripcionesServer) CrearAcceso(w http.ResponseWriter, r *http.Request) {
	var nuevo models.AccesoDigital
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.CorreoAcceso) == "" {
		RespondError(w, http.StatusBadRequest, "el campo correo_acceso es obligatorio")
		return
	}
	if strings.TrimSpace(nuevo.Perfil) == "" {
		RespondError(w, http.StatusBadRequest, "el campo perfil es obligatorio")
		return
	}
	if strings.TrimSpace(nuevo.Estado) == "" {
		RespondError(w, http.StatusBadRequest, "el campo estado es obligatorio")
		return
	}
	if nuevo.SuscripcionClienteID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo suscripcion_cliente_id es obligatorio")
		return
	}
	RespondJSON(w, http.StatusCreated, s.Storage.CrearAcceso(nuevo))
}

func (s *SuscripcionesServer) ActualizarAcceso(w http.ResponseWriter, r *http.Request) {
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
	actualizado, encontrado := s.Storage.ActualizarAcceso(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "acceso digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *SuscripcionesServer) BorrarAcceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if !s.Storage.BorrarAcceso(id) {
		RespondError(w, http.StatusNotFound, "acceso digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
