package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

type mockServicioRepo struct {
	mock.Mock
}

func (m *mockServicioRepo) ListarServicios() []models.ServicioDigital {
	args := m.Called()
	return args.Get(0).([]models.ServicioDigital)
}
func (m *mockServicioRepo) BuscarServicioPorID(id int) (models.ServicioDigital, bool) {
	args := m.Called(id)
	return args.Get(0).(models.ServicioDigital), args.Bool(1)
}
func (m *mockServicioRepo) CrearServicio(s models.ServicioDigital) models.ServicioDigital {
	args := m.Called(s)
	return args.Get(0).(models.ServicioDigital)
}
func (m *mockServicioRepo) ActualizarServicio(id int, datos models.ServicioDigital) (models.ServicioDigital, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.ServicioDigital), args.Bool(1)
}
func (m *mockServicioRepo) BorrarServicio(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func TestServicioDigitalService_Crear_PrecioInvalido(t *testing.T) {
	repo := new(mockServicioRepo)
	repo.On("CrearServicio", mock.Anything).Return(models.ServicioDigital{})
	svc := service.NewServicioDigitalService(repo)

	_, err := svc.Crear(models.ServicioDigital{
		Nombre: "Netflix",
		Precio: 0,
	})

	require.ErrorIs(t, err, service.ErrNombreServicioVacio)
	repo.AssertNotCalled(t, "CrearServicio")
}
