package service

import (
	"strings"

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

// =========================================================
// SUSCRIPCION CLIENTE
// =========================================================

type SuscripcionClienteService struct {
	repo storage.SuscripcionClienteRepository
}

func NewSuscripcionClienteService(repo storage.SuscripcionClienteRepository) *SuscripcionClienteService {
	return &SuscripcionClienteService{repo: repo}
}

func (s *SuscripcionClienteService) Listar() []models.SuscripcionCliente {
	return s.repo.ListarSuscripciones()
}

func (s *SuscripcionClienteService) Obtener(id int) (models.SuscripcionCliente, error) {
	suscripcion, ok := s.repo.BuscarSuscripcionPorID(id)
	if !ok {
		return models.SuscripcionCliente{}, ErrNoEncontrado
	}
	return suscripcion, nil
}

func (s *SuscripcionClienteService) Crear(suscripcion models.SuscripcionCliente) (models.SuscripcionCliente, error) {
	if err := validarSuscripcion(suscripcion); err != nil {
		return models.SuscripcionCliente{}, err
	}
	return s.repo.CrearSuscripcion(suscripcion), nil
}

func (s *SuscripcionClienteService) Actualizar(id int, datos models.SuscripcionCliente) (models.SuscripcionCliente, error) {
	if err := validarSuscripcion(datos); err != nil {
		return models.SuscripcionCliente{}, err
	}
	actualizada, ok := s.repo.ActualizarSuscripcion(id, datos)
	if !ok {
		return models.SuscripcionCliente{}, ErrNoEncontrado
	}
	return actualizada, nil
}

func (s *SuscripcionClienteService) Borrar(id int) error {
	if !s.repo.BorrarSuscripcion(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarSuscripcion(s models.SuscripcionCliente) error {
	if s.ClienteID == 0 {
		return ErrClienteIDRequerido
	}
	if s.ServicioDigitalID == 0 {
		return ErrServicioDigitalIDRequerido
	}
	if strings.TrimSpace(s.FechaInicio) == "" {
		return ErrFechaInicioVacia
	}
	if strings.TrimSpace(s.FechaFin) == "" {
		return ErrFechaFinVacia
	}
	if strings.TrimSpace(s.Estado) == "" {
		return ErrEstadoVacio
	}
	if s.TecnicoID == 0 {
		return ErrTecnicoIDRequerido
	}
	return nil
}

// =========================================================
// ACCESO DIGITAL
// =========================================================

type AccesoDigitalService struct {
	repo storage.AccesoDigitalRepository
}

func NewAccesoDigitalService(repo storage.AccesoDigitalRepository) *AccesoDigitalService {
	return &AccesoDigitalService{repo: repo}
}

func (s *AccesoDigitalService) Listar() []models.AccesoDigital {
	return s.repo.ListarAccesos()
}

func (s *AccesoDigitalService) Obtener(id int) (models.AccesoDigital, error) {
	acceso, ok := s.repo.BuscarAccesoPorID(id)
	if !ok {
		return models.AccesoDigital{}, ErrNoEncontrado
	}
	return acceso, nil
}

func (s *AccesoDigitalService) Crear(acceso models.AccesoDigital) (models.AccesoDigital, error) {
	if err := validarAcceso(acceso); err != nil {
		return models.AccesoDigital{}, err
	}
	return s.repo.CrearAcceso(acceso), nil
}

func (s *AccesoDigitalService) Actualizar(id int, datos models.AccesoDigital) (models.AccesoDigital, error) {
	if err := validarAcceso(datos); err != nil {
		return models.AccesoDigital{}, err
	}
	actualizado, ok := s.repo.ActualizarAcceso(id, datos)
	if !ok {
		return models.AccesoDigital{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *AccesoDigitalService) Borrar(id int) error {
	if !s.repo.BorrarAcceso(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarAcceso(a models.AccesoDigital) error {
	if strings.TrimSpace(a.CorreoAcceso) == "" {
		return ErrCorreoAccesoVacio
	}
	if strings.TrimSpace(a.Perfil) == "" {
		return ErrPerfilVacio
	}
	if strings.TrimSpace(a.Estado) == "" {
		return ErrEstadoVacio
	}
	if a.SuscripcionClienteID == 0 {
		return ErrSuscripcionIDRequerido
	}
	return nil
}
