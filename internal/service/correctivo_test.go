package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

// =========================================================
// MOCK DEL REPOSITORIO
// =========================================================

type mockOrdenRepo struct {
	mock.Mock
}

func (m *mockOrdenRepo) ListarOrdenes() []models.OrdenCorrectiva {
	args := m.Called()
	return args.Get(0).([]models.OrdenCorrectiva)
}
func (m *mockOrdenRepo) BuscarOrdenPorID(id int) (models.OrdenCorrectiva, bool) {
	args := m.Called(id)
	return args.Get(0).(models.OrdenCorrectiva), args.Bool(1)
}
func (m *mockOrdenRepo) CrearOrden(o models.OrdenCorrectiva) models.OrdenCorrectiva {
	args := m.Called(o)
	return args.Get(0).(models.OrdenCorrectiva)
}
func (m *mockOrdenRepo) ActualizarOrden(id int, datos models.OrdenCorrectiva) (models.OrdenCorrectiva, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.OrdenCorrectiva), args.Bool(1)
}
func (m *mockOrdenRepo) ActualizarOrdenParcial(id int, estado string, diagnostico string) (models.OrdenCorrectiva, bool) {
	args := m.Called(id, estado, diagnostico)
	return args.Get(0).(models.OrdenCorrectiva), args.Bool(1)
}
func (m *mockOrdenRepo) BorrarOrden(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// =========================================================
// TEST 1 — SERVICE CON MOCK
// =========================================================

// TestOrdenService_Crear_CodigoVacio prueba que una orden con codigo vacio
// es rechazada y NO llega al repositorio.
func TestOrdenService_Crear_CodigoVacio(t *testing.T) {
	// Preparar
	repo := new(mockOrdenRepo)
	repo.On("CrearOrden", mock.Anything).Return(models.OrdenCorrectiva{})
	svc := service.NewOrdenCorrectivaService(repo)

	// Ejecutar
	_, err := svc.Crear(models.OrdenCorrectiva{
		Codigo:            "",
		ProblemaReportado: "Pantalla rota",
	})

	// Verificar
	require.ErrorIs(t, err, service.ErrCodigoVacio)
	repo.AssertNotCalled(t, "CrearOrden")
}
