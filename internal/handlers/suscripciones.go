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

// SuscripcionesServer guarda el storage que voy a usar en los handlers.
type SuscripcionesServer struct {
	Storage *storage.MemoriaSuscripciones
}

// NewSuscripcionesServer inicializa el server con el storage recibido.
func NewSuscripcionesServer(s *storage.MemoriaSuscripciones) *SuscripcionesServer {
	return &SuscripcionesServer{Storage: s}
}

// SuscripcionesRouter define todas las rutas de mi módulo de suscripciones con subrouter Chi.
func SuscripcionesRouter(store *storage.MemoriaSuscripciones) http.Handler {
	s := NewSuscripcionesServer(store)
	r := chi.NewRouter()

	// Rutas para servicios digitales
	r.Get("/servicios", s.ListarServicios)
	r.Post("/servicios", s.CrearServicio)
	r.Get("/servicios/{id}", s.ObtenerServicio)
	r.Put("/servicios/{id}", s.ActualizarServicio)
	r.Delete("/servicios/{id}", s.BorrarServicio)

	// Rutas para suscripciones de clientes
	r.Get("/suscripciones", s.ListarSuscripciones)
	r.Post("/suscripciones", s.CrearSuscripcion)
	r.Get("/suscripciones/{id}", s.ObtenerSuscripcion)
	r.Put("/suscripciones/{id}", s.ActualizarSuscripcion)
	r.Delete("/suscripciones/{id}", s.BorrarSuscripcion)

	// Rutas para accesos digitales
	r.Get("/accesos", s.ListarAccesos)
	r.Post("/accesos", s.CrearAcceso)
	r.Get("/accesos/{id}", s.ObtenerAcceso)
	r.Put("/accesos/{id}", s.ActualizarAcceso)
	r.Delete("/accesos/{id}", s.BorrarAcceso)

	return r
}

// =========================================================
// SERVICIOS DIGITALES
// =========================================================

// ListarServicios devuelve todos los servicios digitales registrados.
func (s *SuscripcionesServer) ListarServicios(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarServicios())
}

// ObtenerServicio busca un servicio por su ID, si no existe devuelve 404.
func (s *SuscripcionesServer) ObtenerServicio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	servicio, encontrado := s.Storage.BuscarServicioPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "servicio digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, servicio)
}

// CrearServicio recibe los datos del servicio, valida los campos obligatorios y lo guarda.
func (s *SuscripcionesServer) CrearServicio(w http.ResponseWriter, r *http.Request) {
	var nuevo models.ServicioDigital
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	if nuevo.Precio <= 0 {
		RespondError(w, http.StatusBadRequest, "el campo precio debe ser mayor a 0")
		return
	}
	RespondJSON(w, http.StatusCreated, s.Storage.CrearServicio(nuevo))
}

// ActualizarServicio reemplaza completamente un servicio existente por su ID.
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
	if strings.TrimSpace(datos.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}
	actualizado, encontrado := s.Storage.ActualizarServicio(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "servicio digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarServicio elimina un servicio por su ID, si no existe devuelve 404.
func (s *SuscripcionesServer) BorrarServicio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if !s.Storage.BorrarServicio(id) {
		RespondError(w, http.StatusNotFound, "servicio digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// SUSCRIPCIONES CLIENTES
// =========================================================

// ListarSuscripciones devuelve todas las suscripciones de clientes registradas.
func (s *SuscripcionesServer) ListarSuscripciones(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarSuscripciones())
}

// ObtenerSuscripcion busca una suscripción por su ID, si no existe devuelve 404.
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

// CrearSuscripcion recibe los datos de la suscripción, valida los campos obligatorios y la guarda.
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

// ActualizarSuscripcion reemplaza completamente una suscripción existente por su ID.
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
	if datos.ClienteID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo cliente_id es obligatorio")
		return
	}
	actualizada, encontrado := s.Storage.ActualizarSuscripcion(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "suscripción no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

// BorrarSuscripcion elimina una suscripción por su ID, si no existe devuelve 404.
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

// =========================================================
// ACCESOS DIGITALES
// =========================================================

// ListarAccesos devuelve todos los accesos digitales registrados.
func (s *SuscripcionesServer) ListarAccesos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarAccesos())
}

// ObtenerAcceso busca un acceso por su ID, si no existe devuelve 404.
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

// CrearAcceso recibe los datos del acceso, valida los campos obligatorios y lo guarda.
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
	if nuevo.SuscripcionClienteID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo suscripcion_cliente_id es obligatorio")
		return
	}
	RespondJSON(w, http.StatusCreated, s.Storage.CrearAcceso(nuevo))
}

// ActualizarAcceso reemplaza completamente un acceso existente por su ID.
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
	if strings.TrimSpace(datos.CorreoAcceso) == "" {
		RespondError(w, http.StatusBadRequest, "el campo correo_acceso es obligatorio")
		return
	}
	actualizado, encontrado := s.Storage.ActualizarAcceso(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "acceso digital no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarAcceso elimina un acceso por su ID, si no existe devuelve 404.
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
