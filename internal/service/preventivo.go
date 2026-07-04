package service

import (
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
