package storage

import (
	"sync"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
)

// MemoriaCorrectivo es el almacén en memoria para todo el módulo correctivo.
// Tiene slices para órdenes, procesos y evidencias, cada uno con su propio contador de IDs.
type MemoriaCorrectivo struct {
	ordenes     []models.OrdenCorrectiva
	nextOrdenID int

	procesos      []models.ProcesoReparacion
	nextProcesoID int

	evidencias      []models.EvidenciaDanio
	nextEvidenciaID int

	mu sync.Mutex
}

// NuevaMemoriaCorrectivo inicializa el storage con los slices vacíos y los IDs desde 1.
func NuevaMemoriaCorrectivo() *MemoriaCorrectivo {
	return &MemoriaCorrectivo{
		ordenes:         []models.OrdenCorrectiva{},
		nextOrdenID:     1,
		procesos:        []models.ProcesoReparacion{},
		nextProcesoID:   1,
		evidencias:      []models.EvidenciaDanio{},
		nextEvidenciaID: 1,
	}
}

// =========================================================
// ORDENES CORRECTIVAS
// =========================================================

// ListarOrdenes devuelve una copia de todas las órdenes para no exponer el slice interno.
func (m *MemoriaCorrectivo) ListarOrdenes() []models.OrdenCorrectiva {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.OrdenCorrectiva, len(m.ordenes))
	copy(copia, m.ordenes)
	return copia
}

// BuscarOrdenPorID recorre las órdenes y devuelve la que coincida con el ID, o false si no existe.
func (m *MemoriaCorrectivo) BuscarOrdenPorID(id int) (models.OrdenCorrectiva, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, o := range m.ordenes {
		if o.ID == id {
			return o, true
		}
	}
	return models.OrdenCorrectiva{}, false
}

// CrearOrden le asigna un ID a la orden y la agrega al slice.
func (m *MemoriaCorrectivo) CrearOrden(o models.OrdenCorrectiva) models.OrdenCorrectiva {
	m.mu.Lock()
	defer m.mu.Unlock()
	o.ID = m.nextOrdenID
	m.nextOrdenID++
	m.ordenes = append(m.ordenes, o)
	return o
}

// ActualizarOrden reemplaza toda la orden por los datos nuevos, conservando el mismo ID.
func (m *MemoriaCorrectivo) ActualizarOrden(id int, datos models.OrdenCorrectiva) (models.OrdenCorrectiva, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, o := range m.ordenes {
		if o.ID == id {
			datos.ID = id
			m.ordenes[i] = datos
			return datos, true
		}
	}
	return models.OrdenCorrectiva{}, false
}

// ActualizarOrdenParcial solo toca el estado y/o diagnóstico, deja el resto igual.
func (m *MemoriaCorrectivo) ActualizarOrdenParcial(id int, estado string, diagnostico string) (models.OrdenCorrectiva, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, o := range m.ordenes {
		if o.ID == id {
			if estado != "" {
				m.ordenes[i].Estado = estado
			}
			if diagnostico != "" {
				m.ordenes[i].Diagnostico = diagnostico
			}
			return m.ordenes[i], true
		}
	}
	return models.OrdenCorrectiva{}, false
}

// BorrarOrden elimina la orden del slice usando append para cerrar el hueco.
func (m *MemoriaCorrectivo) BorrarOrden(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, o := range m.ordenes {
		if o.ID == id {
			m.ordenes = append(m.ordenes[:i], m.ordenes[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// PROCESOS REPARACION
// =========================================================

// ListarProcesos devuelve una copia de todos los procesos registrados.
func (m *MemoriaCorrectivo) ListarProcesos() []models.ProcesoReparacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.ProcesoReparacion, len(m.procesos))
	copy(copia, m.procesos)
	return copia
}

// BuscarProcesoPorID recorre los procesos y devuelve el que coincida con el ID, o false si no existe.
func (m *MemoriaCorrectivo) BuscarProcesoPorID(id int) (models.ProcesoReparacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.procesos {
		if p.ID == id {
			return p, true
		}
	}
	return models.ProcesoReparacion{}, false
}

// CrearProceso le asigna un ID al proceso y lo agrega al slice.
func (m *MemoriaCorrectivo) CrearProceso(p models.ProcesoReparacion) models.ProcesoReparacion {
	m.mu.Lock()
	defer m.mu.Unlock()
	p.ID = m.nextProcesoID
	m.nextProcesoID++
	m.procesos = append(m.procesos, p)
	return p
}

// ActualizarProceso reemplaza todo el proceso con los datos nuevos, conservando el mismo ID.
func (m *MemoriaCorrectivo) ActualizarProceso(id int, datos models.ProcesoReparacion) (models.ProcesoReparacion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.procesos {
		if p.ID == id {
			datos.ID = id
			m.procesos[i] = datos
			return datos, true
		}
	}
	return models.ProcesoReparacion{}, false
}

// BorrarProceso elimina el proceso del slice usando append para cerrar el hueco.
func (m *MemoriaCorrectivo) BorrarProceso(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.procesos {
		if p.ID == id {
			m.procesos = append(m.procesos[:i], m.procesos[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// EVIDENCIAS DANIO
// =========================================================

// ListarEvidencias devuelve una copia de todas las evidencias registradas.
func (m *MemoriaCorrectivo) ListarEvidencias() []models.EvidenciaDanio {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.EvidenciaDanio, len(m.evidencias))
	copy(copia, m.evidencias)
	return copia
}

// BuscarEvidenciaPorID recorre las evidencias y devuelve la que coincida con el ID, o false si no existe.
func (m *MemoriaCorrectivo) BuscarEvidenciaPorID(id int) (models.EvidenciaDanio, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, e := range m.evidencias {
		if e.ID == id {
			return e, true
		}
	}
	return models.EvidenciaDanio{}, false
}

// CrearEvidencia le asigna un ID a la evidencia y la agrega al slice.
func (m *MemoriaCorrectivo) CrearEvidencia(e models.EvidenciaDanio) models.EvidenciaDanio {
	m.mu.Lock()
	defer m.mu.Unlock()
	e.ID = m.nextEvidenciaID
	m.nextEvidenciaID++
	m.evidencias = append(m.evidencias, e)
	return e
}

// ActualizarEvidencia reemplaza toda la evidencia con los datos nuevos, conservando el mismo ID.
func (m *MemoriaCorrectivo) ActualizarEvidencia(id int, datos models.EvidenciaDanio) (models.EvidenciaDanio, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.evidencias {
		if e.ID == id {
			datos.ID = id
			m.evidencias[i] = datos
			return datos, true
		}
	}
	return models.EvidenciaDanio{}, false
}

// BorrarEvidencia elimina la evidencia del slice usando append para cerrar el hueco.
func (m *MemoriaCorrectivo) BorrarEvidencia(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, e := range m.evidencias {
		if e.ID == id {
			m.evidencias = append(m.evidencias[:i], m.evidencias[i+1:]...)
			return true
		}
	}
	return false
}
