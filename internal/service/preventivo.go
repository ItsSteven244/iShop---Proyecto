package service

import (
	"strings"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

type MantenimientoPreventivoService struct {
	repo storage.MantenimientoPreventivoRepository
}

func NewMantenimientoPreventivoService(repo storage.MantenimientoPreventivoRepository) *MantenimientoPreventivoService {
	return &MantenimientoPreventivoService{repo: repo}
}

func (s *MantenimientoPreventivoService) Listar() []models.MantenimientoPreventivo {
	return s.repo.ListarMantenimientos()
}

func (s *MantenimientoPreventivoService) Obtener(id int) (models.MantenimientoPreventivo, error) {
	mant, ok := s.repo.BuscarMantenimientoPorID(id)
	if !ok {
		return models.MantenimientoPreventivo{}, ErrNoEncontrado
	}
	return mant, nil
}

func (s *MantenimientoPreventivoService) Crear(mant models.MantenimientoPreventivo) (models.MantenimientoPreventivo, error) {
	if err := validarMantenimiento(mant); err != nil {
		return models.MantenimientoPreventivo{}, err
	}
	return s.repo.CrearMantenimiento(mant), nil
}

func (s *MantenimientoPreventivoService) Actualizar(id int, datos models.MantenimientoPreventivo) (models.MantenimientoPreventivo, error) {
	if err := validarMantenimiento(datos); err != nil {
		return models.MantenimientoPreventivo{}, err
	}
	actualizado, ok := s.repo.ActualizarMantenimiento(id, datos)
	if !ok {
		return models.MantenimientoPreventivo{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *MantenimientoPreventivoService) Borrar(id int) error {
	if !s.repo.BorrarMantenimiento(id) {
		return ErrNoEncontrado
	}
	return nil
}

// validarMantenimiento aplica la regla de negocio: equipo y fecha_programada son obligatorios.
func validarMantenimiento(mant models.MantenimientoPreventivo) error {
	if mant.Equipo == "" {
		return ErrEquipoVacio
	}
	if mant.FechaProgramada == "" {
		return ErrFechaProgramadaVacia
	}
	return nil
}

// =========================================================
// TAREA PREVENTIVA.
// =========================================================

type TareaPreventivaService struct {
	repo storage.TareaPreventivaRepository
}

func NewTareaPreventivaService(repo storage.TareaPreventivaRepository) *TareaPreventivaService {
	return &TareaPreventivaService{repo: repo}
}

func (s *TareaPreventivaService) Listar() []models.TareaPreventiva {
	return s.repo.ListarTareas()
}

func (s *TareaPreventivaService) Obtener(id int) (models.TareaPreventiva, error) {
	tarea, ok := s.repo.BuscarTareaPorID(id)
	if !ok {
		return models.TareaPreventiva{}, ErrNoEncontrado
	}
	return tarea, nil
}

func (s *TareaPreventivaService) Crear(tarea models.TareaPreventiva) (models.TareaPreventiva, error) {
	if err := validarTarea(tarea); err != nil {
		return models.TareaPreventiva{}, err
	}
	return s.repo.CrearTarea(tarea), nil
}

func (s *TareaPreventivaService) Actualizar(id int, datos models.TareaPreventiva) (models.TareaPreventiva, error) {
	if err := validarTarea(datos); err != nil {
		return models.TareaPreventiva{}, err
	}
	actualizada, ok := s.repo.ActualizarTarea(id, datos)
	if !ok {
		return models.TareaPreventiva{}, ErrNoEncontrado
	}
	return actualizada, nil
}

func (s *TareaPreventivaService) Borrar(id int) error {
	if !s.repo.BorrarTarea(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarTarea(t models.TareaPreventiva) error {
	if strings.TrimSpace(t.Descripcion) == "" {
		return ErrDescripcionTareaVacia
	}
	if t.MantenimientoPreventivoID == 0 {
		return ErrMantenimientoIDRequerido
	}
	return nil
}

// =========================================================
// INSUMO PREVENTIVO
// =========================================================

type InsumoPreventivoService struct {
	repo storage.InsumoPreventivoRepository
}

func NewInsumoPreventivoService(repo storage.InsumoPreventivoRepository) *InsumoPreventivoService {
	return &InsumoPreventivoService{repo: repo}
}

func (s *InsumoPreventivoService) Listar() []models.InsumoPreventivo {
	return s.repo.ListarInsumos()
}

func (s *InsumoPreventivoService) Obtener(id int) (models.InsumoPreventivo, error) {
	insumo, ok := s.repo.BuscarInsumoPorID(id)
	if !ok {
		return models.InsumoPreventivo{}, ErrNoEncontrado
	}
	return insumo, nil
}

func (s *InsumoPreventivoService) Crear(insumo models.InsumoPreventivo) (models.InsumoPreventivo, error) {
	if err := validarInsumo(insumo); err != nil {
		return models.InsumoPreventivo{}, err
	}
	return s.repo.CrearInsumo(insumo), nil
}

func (s *InsumoPreventivoService) Actualizar(id int, datos models.InsumoPreventivo) (models.InsumoPreventivo, error) {
	if err := validarInsumo(datos); err != nil {
		return models.InsumoPreventivo{}, err
	}
	actualizado, ok := s.repo.ActualizarInsumo(id, datos)
	if !ok {
		return models.InsumoPreventivo{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *InsumoPreventivoService) Borrar(id int) error {
	if !s.repo.BorrarInsumo(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarInsumo(ins models.InsumoPreventivo) error {
	if strings.TrimSpace(ins.Nombre) == "" {
		return ErrNombreInsumoVacio
	}
	if ins.MantenimientoPreventivoID == 0 {
		return ErrMantenimientoIDRequerido
	}
	return nil
}
