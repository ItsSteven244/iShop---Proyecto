package storage

import (
	"gorm.io/gorm"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
)

// =========================================================
// MÓDULO CORRECTIVO - Steven
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
