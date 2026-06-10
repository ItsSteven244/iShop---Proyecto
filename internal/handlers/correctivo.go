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

// CorrectivosServer agrupa el storage del módulo correctivo.
type CorrectivosServer struct {
	Storage *storage.MemoriaCorrectivo
}

// NewCorrectivosServer construye un CorrectivosServer listo para usar.
func NewCorrectivosServer(s *storage.MemoriaCorrectivo) *CorrectivosServer {
	return &CorrectivosServer{Storage: s}
}

// CorrectivosRouter registra todas las rutas del módulo correctivo.
func CorrectivosRouter(store *storage.MemoriaCorrectivo) http.Handler {
	s := NewCorrectivosServer(store)
	r := chi.NewRouter()

	// Ordenes correctivas
	r.Get("/ordenes", s.ListarOrdenes)
	r.Post("/ordenes", s.CrearOrden)
	r.Get("/ordenes/{id}", s.ObtenerOrden)
	r.Put("/ordenes/{id}", s.ActualizarOrden)
	r.Patch("/ordenes/{id}", s.ActualizarEstadoOrden)
	r.Delete("/ordenes/{id}", s.BorrarOrden)

	// Procesos de reparacion
	r.Get("/procesos", s.ListarProcesos)
	r.Post("/procesos", s.CrearProceso)
	r.Get("/procesos/{id}", s.ObtenerProceso)
	r.Put("/procesos/{id}", s.ActualizarProceso)
	r.Delete("/procesos/{id}", s.BorrarProceso)

	// Evidencias de daño
	r.Get("/evidencias", s.ListarEvidencias)
	r.Post("/evidencias", s.CrearEvidencia)
	r.Get("/evidencias/{id}", s.ObtenerEvidencia)
	r.Put("/evidencias/{id}", s.ActualizarEvidencia)
	r.Delete("/evidencias/{id}", s.BorrarEvidencia)

	return r
}

// =========================================================
// ORDENES CORRECTIVAS
// =========================================================

func (s *CorrectivosServer) ListarOrdenes(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarOrdenes())
}

func (s *CorrectivosServer) ObtenerOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	orden, encontrado := s.Storage.BuscarOrdenPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "orden correctiva no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, orden)
}

func (s *CorrectivosServer) CrearOrden(w http.ResponseWriter, r *http.Request) {
	var nueva models.OrdenCorrectiva
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nueva.Codigo) == "" {
		RespondError(w, http.StatusBadRequest, "el campo codigo es obligatorio")
		return
	}
	if strings.TrimSpace(nueva.ProblemaReportado) == "" {
		RespondError(w, http.StatusBadRequest, "el campo problema_reportado es obligatorio")
		return
	}
	RespondJSON(w, http.StatusCreated, s.Storage.CrearOrden(nueva))
}

func (s *CorrectivosServer) ActualizarOrden(w http.ResponseWriter, r *http.Request) {
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
	if strings.TrimSpace(datos.ProblemaReportado) == "" {
		RespondError(w, http.StatusBadRequest, "el campo problema_reportado es obligatorio")
		return
	}
	actualizada, encontrado := s.Storage.ActualizarOrden(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "orden correctiva no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *CorrectivosServer) ActualizarEstadoOrden(w http.ResponseWriter, r *http.Request) {
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
	actualizada, encontrado := s.Storage.ActualizarOrdenParcial(id, body.Estado, body.Diagnostico)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "orden correctiva no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *CorrectivosServer) BorrarOrden(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if !s.Storage.BorrarOrden(id) {
		RespondError(w, http.StatusNotFound, "orden correctiva no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// PROCESOS REPARACION
// =========================================================

func (s *CorrectivosServer) ListarProcesos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarProcesos())
}

func (s *CorrectivosServer) ObtenerProceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	proceso, encontrado := s.Storage.BuscarProcesoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "proceso de reparación no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, proceso)
}

func (s *CorrectivosServer) CrearProceso(w http.ResponseWriter, r *http.Request) {
	var nuevo models.ProcesoReparacion
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nuevo.Etapa) == "" {
		RespondError(w, http.StatusBadRequest, "el campo etapa es obligatorio")
		return
	}
	if nuevo.OrdenCorrectivaID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo orden_correctiva_id es obligatorio")
		return
	}
	RespondJSON(w, http.StatusCreated, s.Storage.CrearProceso(nuevo))
}

func (s *CorrectivosServer) ActualizarProceso(w http.ResponseWriter, r *http.Request) {
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
	if strings.TrimSpace(datos.Etapa) == "" {
		RespondError(w, http.StatusBadRequest, "el campo etapa es obligatorio")
		return
	}
	actualizado, encontrado := s.Storage.ActualizarProceso(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "proceso de reparación no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *CorrectivosServer) BorrarProceso(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if !s.Storage.BorrarProceso(id) {
		RespondError(w, http.StatusNotFound, "proceso de reparación no encontrado")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}

// =========================================================
// EVIDENCIAS DANIO
// =========================================================

func (s *CorrectivosServer) ListarEvidencias(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarEvidencias())
}

func (s *CorrectivosServer) ObtenerEvidencia(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	evidencia, encontrado := s.Storage.BuscarEvidenciaPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "evidencia de daño no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, evidencia)
}

func (s *CorrectivosServer) CrearEvidencia(w http.ResponseWriter, r *http.Request) {
	var nueva models.EvidenciaDanio
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nueva.Descripcion) == "" {
		RespondError(w, http.StatusBadRequest, "el campo descripcion es obligatorio")
		return
	}
	if nueva.OrdenCorrectivaID == 0 {
		RespondError(w, http.StatusBadRequest, "el campo orden_correctiva_id es obligatorio")
		return
	}
	RespondJSON(w, http.StatusCreated, s.Storage.CrearEvidencia(nueva))
}

func (s *CorrectivosServer) ActualizarEvidencia(w http.ResponseWriter, r *http.Request) {
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
	if strings.TrimSpace(datos.Descripcion) == "" {
		RespondError(w, http.StatusBadRequest, "el campo descripcion es obligatorio")
		return
	}
	actualizada, encontrado := s.Storage.ActualizarEvidencia(id, datos)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "evidencia de daño no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *CorrectivosServer) BorrarEvidencia(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if !s.Storage.BorrarEvidencia(id) {
		RespondError(w, http.StatusNotFound, "evidencia de daño no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
