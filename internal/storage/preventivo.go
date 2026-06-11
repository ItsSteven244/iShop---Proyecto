package storage

import (
	"sync"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
)

// MemoriaPreventivo mantiene en un solo lugar todos los datos del módulo preventivo:
// Mantenimientos, Tareas e Insumos.
type MemoriaPreventivo struct {
	mantenimientos      []models.MantenimientoPreventivo
	nextMantenimientoID int

	tareas      []models.TareaPreventiva
	nextTareaID int

	insumos      []models.InsumoPreventivo
	nextInsumoID int

	mu sync.Mutex
}

// NuevaMemoriaPreventivo crea un almacén vacío y listo para usar.
func NuevaMemoriaPreventivo() *MemoriaPreventivo {
	return &MemoriaPreventivo{
		mantenimientos:      []models.MantenimientoPreventivo{},
		nextMantenimientoID: 1,
		tareas:              []models.TareaPreventiva{},
		nextTareaID:         1,
		insumos:             []models.InsumoPreventivo{},
		nextInsumoID:        1,
	}
}

// =========================================================
// MANTENIMIENTOS PREVENTIVOS
// =========================================================

// ListarMantenimientos devuelve todos los mantenimientos en memoria.
func (m *MemoriaPreventivo) ListarMantenimientos() []models.MantenimientoPreventivo {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.MantenimientoPreventivo, len(m.mantenimientos))
	copy(copia, m.mantenimientos)
	return copia
}

// BuscarMantenimientoPorID devuelve el mantenimiento con el ID dado (patrón comma-ok).
func (m *MemoriaPreventivo) BuscarMantenimientoPorID(id int) (models.MantenimientoPreventivo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, mant := range m.mantenimientos {
		if mant.ID == id {
			return mant, true
		}
	}
	return models.MantenimientoPreventivo{}, false
}

// CrearMantenimiento agrega un mantenimiento nuevo y devuelve el mantenimiento con ID asignado.
func (m *MemoriaPreventivo) CrearMantenimiento(mant models.MantenimientoPreventivo) models.MantenimientoPreventivo {
	m.mu.Lock()
	defer m.mu.Unlock()
	mant.ID = m.nextMantenimientoID
	m.nextMantenimientoID++
	m.mantenimientos = append(m.mantenimientos, mant)
	return mant
}

// ActualizarMantenimiento reemplaza el mantenimiento con el ID dado.
func (m *MemoriaPreventivo) ActualizarMantenimiento(id int, datos models.MantenimientoPreventivo) (models.MantenimientoPreventivo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, mant := range m.mantenimientos {
		if mant.ID == id {
			datos.ID = id
			m.mantenimientos[i] = datos
			return datos, true
		}
	}
	return models.MantenimientoPreventivo{}, false
}

// BorrarMantenimiento elimina el mantenimiento con el ID dado.
func (m *MemoriaPreventivo) BorrarMantenimiento(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, mant := range m.mantenimientos {
		if mant.ID == id {
			m.mantenimientos = append(m.mantenimientos[:i], m.mantenimientos[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// TAREAS PREVENTIVAS
// =========================================================

// ListarTareas devuelve todas las tareas en memoria.
func (m *MemoriaPreventivo) ListarTareas() []models.TareaPreventiva {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.TareaPreventiva, len(m.tareas))
	copy(copia, m.tareas)
	return copia
}

// BuscarTareaPorID devuelve la tarea con el ID dado (patrón comma-ok).
func (m *MemoriaPreventivo) BuscarTareaPorID(id int) (models.TareaPreventiva, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, t := range m.tareas {
		if t.ID == id {
			return t, true
		}
	}
	return models.TareaPreventiva{}, false
}

// CrearTarea agrega una tarea nueva y devuelve la tarea con ID asignado.
func (m *MemoriaPreventivo) CrearTarea(t models.TareaPreventiva) models.TareaPreventiva {
	m.mu.Lock()
	defer m.mu.Unlock()
	t.ID = m.nextTareaID
	m.nextTareaID++
	m.tareas = append(m.tareas, t)
	return t
}

// ActualizarTarea reemplaza la tarea con el ID dado.
func (m *MemoriaPreventivo) ActualizarTarea(id int, datos models.TareaPreventiva) (models.TareaPreventiva, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, t := range m.tareas {
		if t.ID == id {
			datos.ID = id
			m.tareas[i] = datos
			return datos, true
		}
	}
	return models.TareaPreventiva{}, false
}

// BorrarTarea elimina la tarea con el ID dado.
func (m *MemoriaPreventivo) BorrarTarea(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, t := range m.tareas {
		if t.ID == id {
			m.tareas = append(m.tareas[:i], m.tareas[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// INSUMOS PREVENTIVOS
// =========================================================

// ListarInsumos devuelve todos los insumos en memoria.
func (m *MemoriaPreventivo) ListarInsumos() []models.InsumoPreventivo {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.InsumoPreventivo, len(m.insumos))
	copy(copia, m.insumos)
	return copia
}

// BuscarInsumoPorID devuelve el insumo con el ID dado (patrón comma-ok).
func (m *MemoriaPreventivo) BuscarInsumoPorID(id int) (models.InsumoPreventivo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, ins := range m.insumos {
		if ins.ID == id {
			return ins, true
		}
	}
	return models.InsumoPreventivo{}, false
}

// CrearInsumo agrega un insumo nuevo y devuelve el insumo con ID asignado.
func (m *MemoriaPreventivo) CrearInsumo(ins models.InsumoPreventivo) models.InsumoPreventivo {
	m.mu.Lock()
	defer m.mu.Unlock()
	ins.ID = m.nextInsumoID
	m.nextInsumoID++
	m.insumos = append(m.insumos, ins)
	return ins
}

// ActualizarInsumo reemplaza el insumo con el ID dado.
func (m *MemoriaPreventivo) ActualizarInsumo(id int, datos models.InsumoPreventivo) (models.InsumoPreventivo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, ins := range m.insumos {
		if ins.ID == id {
			datos.ID = id
			m.insumos[i] = datos
			return datos, true
		}
	}

	return models.InsumoPreventivo{}, false
}

// BorrarInsumo elimina el insumo con el ID dado.
func (m *MemoriaPreventivo) BorrarInsumo(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, ins := range m.insumos {
		if ins.ID == id {
			m.insumos = append(m.insumos[:i], m.insumos[i+1:]...)
			return true
		}
	}
	return false
}
