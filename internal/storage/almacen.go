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
