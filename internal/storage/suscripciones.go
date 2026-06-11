package storage

import (
	"sync"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
)

// MemoriaSuscripciones es el almacén en memoria para todo el módulo de suscripciones.
// Tiene slices para servicios, suscripciones y accesos, cada uno con su propio contador de IDs.
type MemoriaSuscripciones struct {
	servicios      []models.ServicioDigital
	nextServicioID int

	suscripciones     []models.SuscripcionCliente
	nextSuscripcionID int

	accesos      []models.AccesoDigital
	nextAccesoID int

	mu sync.Mutex
}

// NuevaMemoriaSuscripciones inicializa el storage con los slices vacíos y los IDs desde 1.
func NuevaMemoriaSuscripciones() *MemoriaSuscripciones {
	return &MemoriaSuscripciones{
		servicios:         []models.ServicioDigital{},
		nextServicioID:    1,
		suscripciones:     []models.SuscripcionCliente{},
		nextSuscripcionID: 1,
		accesos:           []models.AccesoDigital{},
		nextAccesoID:      1,
	}
}

// =========================================================
// SERVICIOS DIGITALES
// =========================================================

// ListarServicios devuelve una copia de todos los servicios para no exponer el slice interno.
func (m *MemoriaSuscripciones) ListarServicios() []models.ServicioDigital {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.ServicioDigital, len(m.servicios))
	copy(copia, m.servicios)
	return copia
}

// BuscarServicioPorID recorre los servicios y devuelve el que coincida con el ID, o false si no existe.
func (m *MemoriaSuscripciones) BuscarServicioPorID(id int) (models.ServicioDigital, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, s := range m.servicios {
		if s.ID == id {
			return s, true
		}
	}
	return models.ServicioDigital{}, false
}

// CrearServicio le asigna un ID al servicio y lo agrega al slice.
func (m *MemoriaSuscripciones) CrearServicio(s models.ServicioDigital) models.ServicioDigital {
	m.mu.Lock()
	defer m.mu.Unlock()
	s.ID = m.nextServicioID
	m.nextServicioID++
	m.servicios = append(m.servicios, s)
	return s
}

// ActualizarServicio reemplaza todo el servicio con los datos nuevos, conservando el mismo ID.
func (m *MemoriaSuscripciones) ActualizarServicio(id int, datos models.ServicioDigital) (models.ServicioDigital, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.servicios {
		if s.ID == id {
			datos.ID = id
			m.servicios[i] = datos
			return datos, true
		}
	}
	return models.ServicioDigital{}, false
}

// BorrarServicio elimina el servicio del slice usando append para cerrar el hueco.
func (m *MemoriaSuscripciones) BorrarServicio(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.servicios {
		if s.ID == id {
			m.servicios = append(m.servicios[:i], m.servicios[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// SUSCRIPCIONES CLIENTES
// =========================================================

// ListarSuscripciones devuelve una copia de todas las suscripciones registradas.
func (m *MemoriaSuscripciones) ListarSuscripciones() []models.SuscripcionCliente {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.SuscripcionCliente, len(m.suscripciones))
	copy(copia, m.suscripciones)
	return copia
}

// BuscarSuscripcionPorID recorre las suscripciones y devuelve la que coincida con el ID, o false si no existe.
func (m *MemoriaSuscripciones) BuscarSuscripcionPorID(id int) (models.SuscripcionCliente, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, s := range m.suscripciones {
		if s.ID == id {
			return s, true
		}
	}
	return models.SuscripcionCliente{}, false
}

// CrearSuscripcion le asigna un ID a la suscripción y la agrega al slice.
func (m *MemoriaSuscripciones) CrearSuscripcion(s models.SuscripcionCliente) models.SuscripcionCliente {
	m.mu.Lock()
	defer m.mu.Unlock()
	s.ID = m.nextSuscripcionID
	m.nextSuscripcionID++
	m.suscripciones = append(m.suscripciones, s)
	return s
}

// ActualizarSuscripcion reemplaza toda la suscripción con los datos nuevos, conservando el mismo ID.
func (m *MemoriaSuscripciones) ActualizarSuscripcion(id int, datos models.SuscripcionCliente) (models.SuscripcionCliente, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.suscripciones {
		if s.ID == id {
			datos.ID = id
			m.suscripciones[i] = datos
			return datos, true
		}
	}
	return models.SuscripcionCliente{}, false
}

// BorrarSuscripcion elimina la suscripción del slice usando append para cerrar el hueco.
func (m *MemoriaSuscripciones) BorrarSuscripcion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, s := range m.suscripciones {
		if s.ID == id {
			m.suscripciones = append(m.suscripciones[:i], m.suscripciones[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// ACCESOS DIGITALES
// =========================================================

// ListarAccesos devuelve una copia de todos los accesos registrados.
func (m *MemoriaSuscripciones) ListarAccesos() []models.AccesoDigital {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.AccesoDigital, len(m.accesos))
	copy(copia, m.accesos)
	return copia
}

// BuscarAccesoPorID recorre los accesos y devuelve el que coincida con el ID, o false si no existe.
func (m *MemoriaSuscripciones) BuscarAccesoPorID(id int) (models.AccesoDigital, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, a := range m.accesos {
		if a.ID == id {
			return a, true
		}
	}
	return models.AccesoDigital{}, false
}

// CrearAcceso le asigna un ID al acceso y lo agrega al slice.
func (m *MemoriaSuscripciones) CrearAcceso(a models.AccesoDigital) models.AccesoDigital {
	m.mu.Lock()
	defer m.mu.Unlock()
	a.ID = m.nextAccesoID
	m.nextAccesoID++
	m.accesos = append(m.accesos, a)
	return a
}

// ActualizarAcceso reemplaza todo el acceso con los datos nuevos, conservando el mismo ID.
func (m *MemoriaSuscripciones) ActualizarAcceso(id int, datos models.AccesoDigital) (models.AccesoDigital, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, a := range m.accesos {
		if a.ID == id {
			datos.ID = id
			m.accesos[i] = datos
			return datos, true
		}
	}
	return models.AccesoDigital{}, false
}

// BorrarAcceso elimina el acceso del slice usando append para cerrar el hueco.
func (m *MemoriaSuscripciones) BorrarAcceso(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, a := range m.accesos {
		if a.ID == id {
			m.accesos = append(m.accesos[:i], m.accesos[i+1:]...)
			return true
		}
	}
	return false
}
