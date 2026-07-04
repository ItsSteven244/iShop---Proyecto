package storage

import "github.com/ItsSteven244/iShop---Proyecto/internal/models"

// =========================================================
// MÓDULO CORRECTIVO
// =========================================================

type OrdenCorrectivaRepository interface {
	ListarOrdenes() []models.OrdenCorrectiva
	BuscarOrdenPorID(id int) (models.OrdenCorrectiva, bool)
	CrearOrden(o models.OrdenCorrectiva) models.OrdenCorrectiva
	ActualizarOrden(id int, datos models.OrdenCorrectiva) (models.OrdenCorrectiva, bool)
	ActualizarOrdenParcial(id int, estado string, diagnostico string) (models.OrdenCorrectiva, bool)
	BorrarOrden(id int) bool
}

type ProcesoReparacionRepository interface {
	ListarProcesos() []models.ProcesoReparacion
	BuscarProcesoPorID(id int) (models.ProcesoReparacion, bool)
	CrearProceso(p models.ProcesoReparacion) models.ProcesoReparacion
	ActualizarProceso(id int, datos models.ProcesoReparacion) (models.ProcesoReparacion, bool)
	BorrarProceso(id int) bool
}

type EvidenciaDanioRepository interface {
	ListarEvidencias() []models.EvidenciaDanio
	BuscarEvidenciaPorID(id int) (models.EvidenciaDanio, bool)
	CrearEvidencia(e models.EvidenciaDanio) models.EvidenciaDanio
	ActualizarEvidencia(id int, datos models.EvidenciaDanio) (models.EvidenciaDanio, bool)
	BorrarEvidencia(id int) bool
}

type CorrectivoRepository interface {
	OrdenCorrectivaRepository
	ProcesoReparacionRepository
	EvidenciaDanioRepository
}

// Chequeo en tiempo de compilación
var _ CorrectivoRepository = (*MemoriaCorrectivo)(nil)

// =========================================================
// AUTH
// =========================================================

type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

// =========================================================
// MÓDULO SUSCRIPCIONES
// =========================================================

type ServicioDigitalRepository interface {
	ListarServicios() []models.ServicioDigital
	BuscarServicioPorID(id int) (models.ServicioDigital, bool)
	CrearServicio(s models.ServicioDigital) models.ServicioDigital
	ActualizarServicio(id int, datos models.ServicioDigital) (models.ServicioDigital, bool)
	BorrarServicio(id int) bool
}

type SuscripcionClienteRepository interface {
	ListarSuscripciones() []models.SuscripcionCliente
	BuscarSuscripcionPorID(id int) (models.SuscripcionCliente, bool)
	CrearSuscripcion(s models.SuscripcionCliente) models.SuscripcionCliente
	ActualizarSuscripcion(id int, datos models.SuscripcionCliente) (models.SuscripcionCliente, bool)
	BorrarSuscripcion(id int) bool
}

type AccesoDigitalRepository interface {
	ListarAccesos() []models.AccesoDigital
	BuscarAccesoPorID(id int) (models.AccesoDigital, bool)
	CrearAcceso(a models.AccesoDigital) models.AccesoDigital
	ActualizarAcceso(id int, datos models.AccesoDigital) (models.AccesoDigital, bool)
	BorrarAcceso(id int) bool
}

type SuscripcionesRepository interface {
	ServicioDigitalRepository
	SuscripcionClienteRepository
	AccesoDigitalRepository
}

var _ SuscripcionesRepository = (*MemoriaSuscripciones)(nil)
var _ SuscripcionesRepository = (*SuscripcionesGORM)(nil)

// =========================================================
// MÓDULO PREVENTIVO
// =========================================================

type MantenimientoPreventivoRepository interface {
	ListarMantenimientos() []models.MantenimientoPreventivo
	BuscarMantenimientoPorID(id int) (models.MantenimientoPreventivo, bool)
	CrearMantenimiento(mant models.MantenimientoPreventivo) models.MantenimientoPreventivo
	ActualizarMantenimiento(id int, datos models.MantenimientoPreventivo) (models.MantenimientoPreventivo, bool)
	BorrarMantenimiento(id int) bool
}

type TareaPreventivaRepository interface {
	ListarTareas() []models.TareaPreventiva
	BuscarTareaPorID(id int) (models.TareaPreventiva, bool)
	CrearTarea(t models.TareaPreventiva) models.TareaPreventiva
	ActualizarTarea(id int, datos models.TareaPreventiva) (models.TareaPreventiva, bool)
	BorrarTarea(id int) bool
}

type InsumoPreventivoRepository interface {
	ListarInsumos() []models.InsumoPreventivo
	BuscarInsumoPorID(id int) (models.InsumoPreventivo, bool)
	CrearInsumo(ins models.InsumoPreventivo) models.InsumoPreventivo
	ActualizarInsumo(id int, datos models.InsumoPreventivo) (models.InsumoPreventivo, bool)
	BorrarInsumo(id int) bool
}

type PreventivoRepository interface {
	MantenimientoPreventivoRepository
	TareaPreventivaRepository
	InsumoPreventivoRepository
}

var _ PreventivoRepository = (*MemoriaPreventivo)(nil)
var _ PreventivoRepository = (*PreventivoGORM)(nil)
