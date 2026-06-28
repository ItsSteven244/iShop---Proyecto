package service

import (
	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

// =========================================================
// ORDEN CORRECTIVA
// =========================================================

type OrdenCorrectivaService struct {
	repo storage.OrdenCorrectivaRepository
}

func NewOrdenCorrectivaService(repo storage.OrdenCorrectivaRepository) *OrdenCorrectivaService {
	return &OrdenCorrectivaService{repo: repo}
}

func (s *OrdenCorrectivaService) Listar() []models.OrdenCorrectiva {
	return s.repo.ListarOrdenes()
}

func (s *OrdenCorrectivaService) Obtener(id int) (models.OrdenCorrectiva, error) {
	o, ok := s.repo.BuscarOrdenPorID(id)
	if !ok {
		return models.OrdenCorrectiva{}, ErrNoEncontrado
	}
	return o, nil
}

func (s *OrdenCorrectivaService) Crear(o models.OrdenCorrectiva) (models.OrdenCorrectiva, error) {
	if err := validarOrden(o); err != nil {
		return models.OrdenCorrectiva{}, err
	}
	return s.repo.CrearOrden(o), nil
}

func (s *OrdenCorrectivaService) Actualizar(id int, o models.OrdenCorrectiva) (models.OrdenCorrectiva, error) {
	if err := validarOrden(o); err != nil {
		return models.OrdenCorrectiva{}, err
	}
	actualizada, ok := s.repo.ActualizarOrden(id, o)
	if !ok {
		return models.OrdenCorrectiva{}, ErrNoEncontrado
	}
	return actualizada, nil
}

func (s *OrdenCorrectivaService) ActualizarParcial(id int, estado string, diagnostico string) (models.OrdenCorrectiva, error) {
	actualizada, ok := s.repo.ActualizarOrdenParcial(id, estado, diagnostico)
	if !ok {
		return models.OrdenCorrectiva{}, ErrNoEncontrado
	}
	return actualizada, nil
}

func (s *OrdenCorrectivaService) Borrar(id int) error {
	if !s.repo.BorrarOrden(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarOrden(o models.OrdenCorrectiva) error {
	if o.Codigo == "" {
		return ErrCodigoVacio
	}
	if o.ProblemaReportado == "" {
		return ErrProblemaVacio
	}
	return nil
}

// =========================================================
// PROCESO REPARACION
// =========================================================

type ProcesoReparacionService struct {
	repo storage.ProcesoReparacionRepository
}

func NewProcesoReparacionService(repo storage.ProcesoReparacionRepository) *ProcesoReparacionService {
	return &ProcesoReparacionService{repo: repo}
}

func (s *ProcesoReparacionService) Listar() []models.ProcesoReparacion {
	return s.repo.ListarProcesos()
}

func (s *ProcesoReparacionService) Obtener(id int) (models.ProcesoReparacion, error) {
	p, ok := s.repo.BuscarProcesoPorID(id)
	if !ok {
		return models.ProcesoReparacion{}, ErrNoEncontrado
	}
	return p, nil
}

func (s *ProcesoReparacionService) Crear(p models.ProcesoReparacion) (models.ProcesoReparacion, error) {
	if err := validarProceso(p); err != nil {
		return models.ProcesoReparacion{}, err
	}
	return s.repo.CrearProceso(p), nil
}

func (s *ProcesoReparacionService) Actualizar(id int, p models.ProcesoReparacion) (models.ProcesoReparacion, error) {
	if err := validarProceso(p); err != nil {
		return models.ProcesoReparacion{}, err
	}
	actualizado, ok := s.repo.ActualizarProceso(id, p)
	if !ok {
		return models.ProcesoReparacion{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *ProcesoReparacionService) Borrar(id int) error {
	if !s.repo.BorrarProceso(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarProceso(p models.ProcesoReparacion) error {
	if p.Etapa == "" {
		return ErrEtapaVacia
	}
	if p.OrdenCorrectivaID == 0 {
		return ErrOrdenIDRequerido
	}
	return nil
}

// =========================================================
// EVIDENCIA DANIO
// =========================================================

type EvidenciaDanioService struct {
	repo storage.EvidenciaDanioRepository
}

func NewEvidenciaDanioService(repo storage.EvidenciaDanioRepository) *EvidenciaDanioService {
	return &EvidenciaDanioService{repo: repo}
}

func (s *EvidenciaDanioService) Listar() []models.EvidenciaDanio {
	return s.repo.ListarEvidencias()
}

func (s *EvidenciaDanioService) Obtener(id int) (models.EvidenciaDanio, error) {
	e, ok := s.repo.BuscarEvidenciaPorID(id)
	if !ok {
		return models.EvidenciaDanio{}, ErrNoEncontrado
	}
	return e, nil
}

func (s *EvidenciaDanioService) Crear(e models.EvidenciaDanio) (models.EvidenciaDanio, error) {
	if err := validarEvidencia(e); err != nil {
		return models.EvidenciaDanio{}, err
	}
	return s.repo.CrearEvidencia(e), nil
}

func (s *EvidenciaDanioService) Actualizar(id int, e models.EvidenciaDanio) (models.EvidenciaDanio, error) {
	if err := validarEvidencia(e); err != nil {
		return models.EvidenciaDanio{}, err
	}
	actualizada, ok := s.repo.ActualizarEvidencia(id, e)
	if !ok {
		return models.EvidenciaDanio{}, ErrNoEncontrado
	}
	return actualizada, nil
}

func (s *EvidenciaDanioService) Borrar(id int) error {
	if !s.repo.BorrarEvidencia(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarEvidencia(e models.EvidenciaDanio) error {
	if e.Descripcion == "" {
		return ErrDescripcionVacia
	}
	if e.OrdenCorrectivaID == 0 {
		return ErrOrdenIDRequerido
	}
	return nil
}
