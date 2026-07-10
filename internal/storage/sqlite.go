package storage

import (
	"gorm.io/gorm"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
)

// =========================================================
// MÓDULO CORRECTIVO - Steven.
// =========================================================

// CorrectivoGORM es la implementación del repositorio correctivo usando GORM + SQLite.
type CorrectivoGORM struct {
	db *gorm.DB
}

// NuevoCorrectivoGORM crea un nuevo repositorio GORM para el módulo correctivo.
func NuevoCorrectivoGORM(db *gorm.DB) *CorrectivoGORM {
	return &CorrectivoGORM{db: db}
}

// Chequeo en tiempo de compilación
var _ CorrectivoRepository = (*CorrectivoGORM)(nil)

// — Órdenes Correctivas —

func (g *CorrectivoGORM) ListarOrdenes() []models.OrdenCorrectiva {
	var ordenes []models.OrdenCorrectiva
	g.db.Find(&ordenes)
	return ordenes
}

func (g *CorrectivoGORM) BuscarOrdenPorID(id int) (models.OrdenCorrectiva, bool) {
	var orden models.OrdenCorrectiva
	if g.db.First(&orden, id).Error != nil {
		return models.OrdenCorrectiva{}, false
	}
	return orden, true
}

func (g *CorrectivoGORM) CrearOrden(o models.OrdenCorrectiva) models.OrdenCorrectiva {
	g.db.Create(&o)
	return o
}

func (g *CorrectivoGORM) ActualizarOrden(id int, datos models.OrdenCorrectiva) (models.OrdenCorrectiva, bool) {
	var orden models.OrdenCorrectiva
	if g.db.First(&orden, id).Error != nil {
		return models.OrdenCorrectiva{}, false
	}
	datos.ID = id
	g.db.Save(&datos)
	return datos, true
}

func (g *CorrectivoGORM) ActualizarOrdenParcial(id int, estado string, diagnostico string) (models.OrdenCorrectiva, bool) {
	var orden models.OrdenCorrectiva
	if g.db.First(&orden, id).Error != nil {
		return models.OrdenCorrectiva{}, false
	}
	if estado != "" {
		orden.Estado = estado
	}
	if diagnostico != "" {
		orden.Diagnostico = diagnostico
	}
	g.db.Save(&orden)
	return orden, true
}

func (g *CorrectivoGORM) BorrarOrden(id int) bool {
	result := g.db.Delete(&models.OrdenCorrectiva{}, id)
	return result.RowsAffected > 0
}

// — Procesos de Reparación —

func (g *CorrectivoGORM) ListarProcesos() []models.ProcesoReparacion {
	var procesos []models.ProcesoReparacion
	g.db.Find(&procesos)
	return procesos
}

func (g *CorrectivoGORM) BuscarProcesoPorID(id int) (models.ProcesoReparacion, bool) {
	var proceso models.ProcesoReparacion
	if g.db.First(&proceso, id).Error != nil {
		return models.ProcesoReparacion{}, false
	}
	return proceso, true
}

func (g *CorrectivoGORM) CrearProceso(p models.ProcesoReparacion) models.ProcesoReparacion {
	g.db.Create(&p)
	return p
}

func (g *CorrectivoGORM) ActualizarProceso(id int, datos models.ProcesoReparacion) (models.ProcesoReparacion, bool) {
	var proceso models.ProcesoReparacion
	if g.db.First(&proceso, id).Error != nil {
		return models.ProcesoReparacion{}, false
	}
	datos.ID = id
	g.db.Save(&datos)
	return datos, true
}

func (g *CorrectivoGORM) BorrarProceso(id int) bool {
	result := g.db.Delete(&models.ProcesoReparacion{}, id)
	return result.RowsAffected > 0
}

// — Evidencias de Daño —

func (g *CorrectivoGORM) ListarEvidencias() []models.EvidenciaDanio {
	var evidencias []models.EvidenciaDanio
	g.db.Find(&evidencias)
	return evidencias
}

func (g *CorrectivoGORM) BuscarEvidenciaPorID(id int) (models.EvidenciaDanio, bool) {
	var evidencia models.EvidenciaDanio
	if g.db.First(&evidencia, id).Error != nil {
		return models.EvidenciaDanio{}, false
	}
	return evidencia, true
}

func (g *CorrectivoGORM) CrearEvidencia(e models.EvidenciaDanio) models.EvidenciaDanio {
	g.db.Create(&e)
	return e
}

func (g *CorrectivoGORM) ActualizarEvidencia(id int, datos models.EvidenciaDanio) (models.EvidenciaDanio, bool) {
	var evidencia models.EvidenciaDanio
	if g.db.First(&evidencia, id).Error != nil {
		return models.EvidenciaDanio{}, false
	}
	datos.ID = id
	g.db.Save(&datos)
	return datos, true
}

func (g *CorrectivoGORM) BorrarEvidencia(id int) bool {
	result := g.db.Delete(&models.EvidenciaDanio{}, id)
	return result.RowsAffected > 0
}

// =========================================================
// MÓDULO SUSCRIPCIONES - Luis
// =========================================================

type SuscripcionesGORM struct {
	db *gorm.DB
}

func NuevoSuscripcionesGORM(db *gorm.DB) *SuscripcionesGORM {
	return &SuscripcionesGORM{db: db}
}

// — Servicios Digitales —

func (g *SuscripcionesGORM) ListarServicios() []models.ServicioDigital {
	var servicios []models.ServicioDigital
	g.db.Find(&servicios)
	return servicios
}

func (g *SuscripcionesGORM) BuscarServicioPorID(id int) (models.ServicioDigital, bool) {
	var servicio models.ServicioDigital
	if g.db.First(&servicio, id).Error != nil {
		return models.ServicioDigital{}, false
	}
	return servicio, true
}

func (g *SuscripcionesGORM) CrearServicio(s models.ServicioDigital) models.ServicioDigital {
	g.db.Create(&s)
	return s
}

func (g *SuscripcionesGORM) ActualizarServicio(id int, datos models.ServicioDigital) (models.ServicioDigital, bool) {
	var servicio models.ServicioDigital
	if g.db.First(&servicio, id).Error != nil {
		return models.ServicioDigital{}, false
	}
	datos.ID = id
	g.db.Save(&datos)
	return datos, true
}

func (g *SuscripcionesGORM) BorrarServicio(id int) bool {
	result := g.db.Delete(&models.ServicioDigital{}, id)
	return result.RowsAffected > 0
}

// — Suscripciones de Clientes —

func (g *SuscripcionesGORM) ListarSuscripciones() []models.SuscripcionCliente {
	var suscripciones []models.SuscripcionCliente
	g.db.Find(&suscripciones)
	return suscripciones
}

func (g *SuscripcionesGORM) BuscarSuscripcionPorID(id int) (models.SuscripcionCliente, bool) {
	var suscripcion models.SuscripcionCliente
	if g.db.First(&suscripcion, id).Error != nil {
		return models.SuscripcionCliente{}, false
	}
	return suscripcion, true
}

func (g *SuscripcionesGORM) CrearSuscripcion(s models.SuscripcionCliente) models.SuscripcionCliente {
	g.db.Create(&s)
	return s
}

func (g *SuscripcionesGORM) ActualizarSuscripcion(id int, datos models.SuscripcionCliente) (models.SuscripcionCliente, bool) {
	var suscripcion models.SuscripcionCliente
	if g.db.First(&suscripcion, id).Error != nil {
		return models.SuscripcionCliente{}, false
	}
	datos.ID = id
	g.db.Save(&datos)
	return datos, true
}

func (g *SuscripcionesGORM) BorrarSuscripcion(id int) bool {
	result := g.db.Delete(&models.SuscripcionCliente{}, id)
	return result.RowsAffected > 0
}

// — Accesos Digitales —

func (g *SuscripcionesGORM) ListarAccesos() []models.AccesoDigital {
	var accesos []models.AccesoDigital
	g.db.Find(&accesos)
	return accesos
}

func (g *SuscripcionesGORM) BuscarAccesoPorID(id int) (models.AccesoDigital, bool) {
	var acceso models.AccesoDigital
	if g.db.First(&acceso, id).Error != nil {
		return models.AccesoDigital{}, false
	}
	return acceso, true
}

func (g *SuscripcionesGORM) CrearAcceso(a models.AccesoDigital) models.AccesoDigital {
	g.db.Create(&a)
	return a
}

func (g *SuscripcionesGORM) ActualizarAcceso(id int, datos models.AccesoDigital) (models.AccesoDigital, bool) {
	var acceso models.AccesoDigital
	if g.db.First(&acceso, id).Error != nil {
		return models.AccesoDigital{}, false
	}
	datos.ID = id
	g.db.Save(&datos)
	return datos, true
}

func (g *SuscripcionesGORM) BorrarAcceso(id int) bool {
	result := g.db.Delete(&models.AccesoDigital{}, id)
	return result.RowsAffected > 0
}

// =========================================================
// MÓDULO PREVENTIVO
// =========================================================

type PreventivoGORM struct {
	db *gorm.DB
}

func NuevoPreventivoGORM(db *gorm.DB) *PreventivoGORM {
	return &PreventivoGORM{db: db}
}

// — Mantenimientos Preventivos —

func (g *PreventivoGORM) ListarMantenimientos() []models.MantenimientoPreventivo {
	var mantenimientos []models.MantenimientoPreventivo
	g.db.Find(&mantenimientos)
	return mantenimientos
}

func (g *PreventivoGORM) BuscarMantenimientoPorID(id int) (models.MantenimientoPreventivo, bool) {
	var mant models.MantenimientoPreventivo
	if g.db.First(&mant, id).Error != nil {
		return models.MantenimientoPreventivo{}, false
	}
	return mant, true
}

func (g *PreventivoGORM) CrearMantenimiento(mant models.MantenimientoPreventivo) models.MantenimientoPreventivo {
	g.db.Create(&mant)
	return mant
}

func (g *PreventivoGORM) ActualizarMantenimiento(id int, datos models.MantenimientoPreventivo) (models.MantenimientoPreventivo, bool) {
	var mant models.MantenimientoPreventivo
	if g.db.First(&mant, id).Error != nil {
		return models.MantenimientoPreventivo{}, false
	}
	datos.ID = id
	g.db.Save(&datos)
	return datos, true
}

func (g *PreventivoGORM) BorrarMantenimiento(id int) bool {
	result := g.db.Delete(&models.MantenimientoPreventivo{}, id)
	return result.RowsAffected > 0
}

// — Tareas Preventivas —

func (g *PreventivoGORM) ListarTareas() []models.TareaPreventiva {
	var tareas []models.TareaPreventiva
	g.db.Find(&tareas)
	return tareas
}

func (g *PreventivoGORM) BuscarTareaPorID(id int) (models.TareaPreventiva, bool) {
	var tarea models.TareaPreventiva
	if g.db.First(&tarea, id).Error != nil {
		return models.TareaPreventiva{}, false
	}
	return tarea, true
}

func (g *PreventivoGORM) CrearTarea(t models.TareaPreventiva) models.TareaPreventiva {
	g.db.Create(&t)
	return t
}

func (g *PreventivoGORM) ActualizarTarea(id int, datos models.TareaPreventiva) (models.TareaPreventiva, bool) {
	var tarea models.TareaPreventiva
	if g.db.First(&tarea, id).Error != nil {
		return models.TareaPreventiva{}, false
	}
	datos.ID = id
	g.db.Save(&datos)
	return datos, true
}

func (g *PreventivoGORM) BorrarTarea(id int) bool {
	result := g.db.Delete(&models.TareaPreventiva{}, id)
	return result.RowsAffected > 0
}

// — Insumos Preventivos —

func (g *PreventivoGORM) ListarInsumos() []models.InsumoPreventivo {
	var insumos []models.InsumoPreventivo
	g.db.Find(&insumos)
	return insumos
}

func (g *PreventivoGORM) BuscarInsumoPorID(id int) (models.InsumoPreventivo, bool) {
	var insumo models.InsumoPreventivo
	if g.db.First(&insumo, id).Error != nil {
		return models.InsumoPreventivo{}, false
	}
	return insumo, true
}

func (g *PreventivoGORM) CrearInsumo(ins models.InsumoPreventivo) models.InsumoPreventivo {
	g.db.Create(&ins)
	return ins
}

func (g *PreventivoGORM) ActualizarInsumo(id int, datos models.InsumoPreventivo) (models.InsumoPreventivo, bool) {
	var insumo models.InsumoPreventivo
	if g.db.First(&insumo, id).Error != nil {
		return models.InsumoPreventivo{}, false
	}
	datos.ID = id
	g.db.Save(&datos)
	return datos, true
}

func (g *PreventivoGORM) BorrarInsumo(id int) bool {
	result := g.db.Delete(&models.InsumoPreventivo{}, id)
	return result.RowsAffected > 0
}
