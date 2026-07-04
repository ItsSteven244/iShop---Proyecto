package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

type mockMantenimientoRepo struct {
	mock.Mock
}

func (m *mockMantenimientoRepo) ListarMantenimientos() []models.MantenimientoPreventivo {
	args := m.Called()
	return args.Get(0).([]models.MantenimientoPreventivo)
}
func (m *mockMantenimientoRepo) BuscarMantenimientoPorID(id int) (models.MantenimientoPreventivo, bool) {
	args := m.Called(id)
	return args.Get(0).(models.MantenimientoPreventivo), args.Bool(1)
}
func (m *mockMantenimientoRepo) CrearMantenimiento(mant models.MantenimientoPreventivo) models.MantenimientoPreventivo {
	args := m.Called(mant)
	return args.Get(0).(models.MantenimientoPreventivo)
}
func (m *mockMantenimientoRepo) ActualizarMantenimiento(id int, datos models.MantenimientoPreventivo) (models.MantenimientoPreventivo, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.MantenimientoPreventivo), args.Bool(1)
}
func (m *mockMantenimientoRepo) BorrarMantenimiento(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// TestMantenimientoPreventivoService_Crear_EquipoVacio prueba que un mantenimiento
// sin equipo es rechazado y NO llega al repositorio.
func TestMantenimientoPreventivoService_Crear_EquipoVacio(t *testing.T) {
	repo := new(mockMantenimientoRepo)
	repo.On("CrearMantenimiento", mock.Anything).Return(models.MantenimientoPreventivo{})
	svc := service.NewMantenimientoPreventivoService(repo)

	_, err := svc.Crear(models.MantenimientoPreventivo{
		FechaProgramada:   "2026-07-01",
		TipoMantenimiento: "Preventivo",
		Equipo:            "",
	})

	require.ErrorIs(t, err, service.ErrEquipoVacio)
	repo.AssertNotCalled(t, "CrearMantenimiento")
}
