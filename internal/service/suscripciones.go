package service

import (
	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

type ServicioDigitalService struct {
	repo storage.ServicioDigitalRepository
}

func NewServicioDigitalService(repo storage.ServicioDigitalRepository) *ServicioDigitalService {
	return &ServicioDigitalService{repo: repo}
}

func (s *ServicioDigitalService) Listar() []models.ServicioDigital {
	return s.repo.ListarServicios()
}

func (s *ServicioDigitalService) Obtener(id int) (models.ServicioDigital, error) {
	servicio, ok := s.repo.BuscarServicioPorID(id)
	if !ok {
		return models.ServicioDigital{}, ErrNoEncontrado
	}
	return servicio, nil
}

func (s *ServicioDigitalService) Crear(servicio models.ServicioDigital) (models.ServicioDigital, error) {
	if err := validarServicio(servicio); err != nil {
		return models.ServicioDigital{}, err
	}
	return s.repo.CrearServicio(servicio), nil
}

func (s *ServicioDigitalService) Actualizar(id int, datos models.ServicioDigital) (models.ServicioDigital, error) {
	if err := validarServicio(datos); err != nil {
		return models.ServicioDigital{}, err
	}
	actualizado, ok := s.repo.ActualizarServicio(id, datos)
	if !ok {
		return models.ServicioDigital{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *ServicioDigitalService) Borrar(id int) error {
	if !s.repo.BorrarServicio(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarServicio(s models.ServicioDigital) error {
	if s.Nombre == "" {
		return ErrNombreServicioVacio
	}
	if s.Precio <= 0 {
		return ErrPrecioInvalido
	}
	return nil
}
