package handlers

import (
	"encoding/json" //traductor go a json
	"net/http"      //para manejar peticiones y respuestas HTTP
	"strconv"       // para convertir texto a número (ej: "3" → 3)
	"strings"       // para manipular texto (ej: quitar espacios)

	"github.com/go-chi/chi/v5" // el router chi — maneja las rutas URL

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

// CorrectivosServer carga el storage que voy a usar en los handlers.
type CorrectivosServer struct {
	Storage *storage.MemoriaCorrectivo
}

// NewCorrectivosServer inicializa el server con el storage recibido.
func NewCorrectivosServer(s *storage.MemoriaCorrectivo) *CorrectivosServer {
	return &CorrectivosServer{Storage: s}
}

// CorrectivosRouter define todas las rutas de mi módulo correctivo con subrouter Chi.
func CorrectivosRouter(store *storage.MemoriaCorrectivo) http.Handler {
	s := NewCorrectivosServer(store)
	r := chi.NewRouter()

	// Rutas para ordenes correctivas
	r.Get("/ordenes", s.ListarOrdenes)
	r.Post("/ordenes", s.CrearOrden)
	r.Get("/ordenes/{id}", s.ObtenerOrden)
	r.Put("/ordenes/{id}", s.ActualizarOrden)
	r.Patch("/ordenes/{id}", s.ActualizarEstadoOrden)
	r.Delete("/ordenes/{id}", s.BorrarOrden)

	// Rutas para procesos de reparacion
	r.Get("/procesos", s.ListarProcesos)
	r.Post("/procesos", s.CrearProceso)
	r.Get("/procesos/{id}", s.ObtenerProceso)
	r.Put("/procesos/{id}", s.ActualizarProceso)
	r.Delete("/procesos/{id}", s.BorrarProceso)

	// Rutas para evidencias de daño
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

// ListarOrdenes devuelve todas las ordenes correctivas registradas.
func (s *CorrectivosServer) ListarOrdenes(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarOrdenes())
}

// ObtenerOrden busca una orden por su ID, si no existe devuelve 404.
func (s *CorrectivosServer) ObtenerOrden(w http.ResponseWriter, r *http.Request) { //w para responder
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int   // y r contiene todo lo que mando el cliente
	if err != nil {                                //si algo salio mal se ejcuta el error de abajo
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	orden, encontrado := s.Storage.BuscarOrdenPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "orden correctiva no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, orden) //si todo sale bien obtiene orden
}

// CrearOrden recibe los datos de la orden, valida los campos obligatorios y la guarda.
func (s *CorrectivosServer) CrearOrden(w http.ResponseWriter, r *http.Request) { //w para responder y r contiene todo lo que mando el cliente
	var nueva models.OrdenCorrectiva                               //creo una nueva orden donde se guardar los datos que envie el cliente
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil { //traduce json y llena el formulario, en caso de error
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error()) //como por ejemplo se envio un json mal envie error 404
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
	RespondJSON(w, http.StatusCreated, s.Storage.CrearOrden(nueva)) //si todo sale bien crea orden
}

// ActualizarOrden reemplaza completamente una orden existente por su ID.
func (s *CorrectivosServer) ActualizarOrden(w http.ResponseWriter, r *http.Request) { //w para responder
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int   // y r contiene todo lo que mando el cliente
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.OrdenCorrectiva
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil { //traduce json y llena el formulario en caso de error
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error()) //como por ejemplo se envio un json mal envie error 404
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
	RespondJSON(w, http.StatusOK, actualizada) //si todo sale bien actualiza
}

// ActualizarEstadoOrden actualiza solo el estado y/o diagnostico de una orden.
func (s *CorrectivosServer) ActualizarEstadoOrden(w http.ResponseWriter, r *http.Request) { //w para responder
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int   // y r contiene todo lo que mando el cliente
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var body struct { //crear formulario en el momento con solo los campos que senecesita, no se necesita todo completo
		Estado      string `json:"estado"`
		Diagnostico string `json:"diagnostico"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil { //traduce json y llena el formulario en caso de error
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error()) //como por ejemplo se envio un json mal envie error 404
		return
	}
	actualizada, encontrado := s.Storage.ActualizarOrdenParcial(id, body.Estado, body.Diagnostico)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "orden correctiva no encontrada")
		return
	}
	RespondJSON(w, http.StatusOK, actualizada) //si todo sale bien actualiza
}

// BorrarOrden elimina una orden por su ID, si no existe devuelve 404.
func (s *CorrectivosServer) BorrarOrden(w http.ResponseWriter, r *http.Request) { //w para responder
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int   // y r contiene todo lo que mando el cliente
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	if !s.Storage.BorrarOrden(id) {
		RespondError(w, http.StatusNotFound, "orden correctiva no encontrada")
		return
	}
	RespondJSON(w, http.StatusNoContent, nil) //si todo salio bien elimina
}

// =========================================================
// PROCESOS REPARACION
// =========================================================

// ListarProcesos devuelve todos los procesos de reparacion registrados.
func (s *CorrectivosServer) ListarProcesos(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarProcesos())
}

// ObtenerProceso busca un proceso por su ID, si no existe devuelve 404.
func (s *CorrectivosServer) ObtenerProceso(w http.ResponseWriter, r *http.Request) { //w para responder
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int   // y r contiene todo lo que mando el cliente
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	proceso, encontrado := s.Storage.BuscarProcesoPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "proceso de reparación no encontrado")
		return
	}
	RespondJSON(w, http.StatusOK, proceso) //si todo sale bien obtiene
}

// CrearProceso recibe los datos del proceso, valida los campos obligatorios y lo guarda.
func (s *CorrectivosServer) CrearProceso(w http.ResponseWriter, r *http.Request) { //w para responder y r contiene todo lo que mando el cliente
	var nuevo models.ProcesoReparacion
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil { //traduce json y llena el formulario en caso de error
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error()) //como por ejemplo se envio un json mal envie error 404
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

// ActualizarProceso reemplaza completamente un proceso existente por su ID.
func (s *CorrectivosServer) ActualizarProceso(w http.ResponseWriter, r *http.Request) { //w para responder y r contiene todo lo que mando el cliente
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.ProcesoReparacion
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil { //traduce json y llena el formulario en caso de error
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error()) //como por ejemplo se envio un json mal envie error 404
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

// BorrarProceso elimina un proceso por su ID, si no existe devuelve 404.
func (s *CorrectivosServer) BorrarProceso(w http.ResponseWriter, r *http.Request) { //w para responder y r contiene todo lo que mando el cliente
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int
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

// ListarEvidencias devuelve todas las evidencias de daño registradas.
func (s *CorrectivosServer) ListarEvidencias(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, s.Storage.ListarEvidencias())
}

// ObtenerEvidencia busca una evidencia por su ID, si no existe devuelve 404.
func (s *CorrectivosServer) ObtenerEvidencia(w http.ResponseWriter, r *http.Request) { //w para responder y r contiene todo lo que mando el cliente
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int
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

// CrearEvidencia recibe los datos de la evidencia, valida los campos obligatorios y la guarda.
func (s *CorrectivosServer) CrearEvidencia(w http.ResponseWriter, r *http.Request) { //w para responder y r contiene todo lo que mando el cliente
	var nueva models.EvidenciaDanio
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil { //traduce json y llena el formulario en caso de error
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error()) //como por ejemplo se envio un json mal envie error 404
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

// ActualizarEvidencia reemplaza completamente una evidencia existente por su ID.
func (s *CorrectivosServer) ActualizarEvidencia(w http.ResponseWriter, r *http.Request) { //w para responder y r contiene todo lo que mando el cliente
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}
	var datos models.EvidenciaDanio
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil { //traduce json y llena el formulario en caso de error
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error()) //como por ejemplo se envio un json mal envie error 404
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

// BorrarEvidencia elimina una evidencia por su ID, si no existe devuelve 404.
func (s *CorrectivosServer) BorrarEvidencia(w http.ResponseWriter, r *http.Request) { //w para responder y r contiene todo lo que mando el cliente
	id, err := strconv.Atoi(chi.URLParam(r, "id")) //atoi convierte string a int
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
