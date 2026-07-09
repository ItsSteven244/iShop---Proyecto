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

// =========================================================
// TEST EXTRAS PARA COBERTURA
// TEST 2 — CREAR: caso de éxito
// =========================================================

func TestOrdenService_Crear_Exito(t *testing.T) {
	repo := new(mockOrdenRepo)
	entrada := models.OrdenCorrectiva{Codigo: "ORD-001", ProblemaReportado: "Pantalla rota"}
	esperada := models.OrdenCorrectiva{ID: 1, Codigo: "ORD-001", ProblemaReportado: "Pantalla rota"}

	repo.On("CrearOrden", entrada).Return(esperada)
	svc := service.NewOrdenCorrectivaService(repo)

	resultado, err := svc.Crear(entrada)

	require.NoError(t, err)
	require.Equal(t, esperada, resultado)
	repo.AssertExpectations(t)
}

// =========================================================
// TEST 3 y 4 — OBTENER: encontrada y no encontrada
// =========================================================

func TestOrdenService_Obtener_Encontrada(t *testing.T) {
	repo := new(mockOrdenRepo)
	esperada := models.OrdenCorrectiva{ID: 1, Codigo: "ORD-001"}
	repo.On("BuscarOrdenPorID", 1).Return(esperada, true)
	svc := service.NewOrdenCorrectivaService(repo)

	resultado, err := svc.Obtener(1)

	require.NoError(t, err)
	require.Equal(t, esperada, resultado)
}

func TestOrdenService_Obtener_NoEncontrada(t *testing.T) {
	repo := new(mockOrdenRepo)
	repo.On("BuscarOrdenPorID", 999).Return(models.OrdenCorrectiva{}, false)
	svc := service.NewOrdenCorrectivaService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

// =========================================================
// TEST 5 — ACTUALIZAR: caso de éxito
// =========================================================

func TestOrdenService_Actualizar_Exito(t *testing.T) {
	repo := new(mockOrdenRepo)
	datos := models.OrdenCorrectiva{Codigo: "ORD-001", ProblemaReportado: "Pantalla rota"}
	actualizada := models.OrdenCorrectiva{ID: 1, Codigo: "ORD-001", ProblemaReportado: "Pantalla rota"}

	repo.On("ActualizarOrden", 1, datos).Return(actualizada, true)
	svc := service.NewOrdenCorrectivaService(repo)

	resultado, err := svc.Actualizar(1, datos)

	require.NoError(t, err)
	require.Equal(t, actualizada, resultado)
}

// =========================================================
// TEST 6 — ACTUALIZAR PARCIAL: caso de éxito
// =========================================================

func TestOrdenService_ActualizarParcial_Exito(t *testing.T) {
	repo := new(mockOrdenRepo)
	actualizada := models.OrdenCorrectiva{ID: 1, Estado: "En proceso", Diagnostico: "Pantalla dañada"}
	repo.On("ActualizarOrdenParcial", 1, "En proceso", "Pantalla dañada").Return(actualizada, true)
	svc := service.NewOrdenCorrectivaService(repo)

	resultado, err := svc.ActualizarParcial(1, "En proceso", "Pantalla dañada")

	require.NoError(t, err)
	require.Equal(t, actualizada, resultado)
}

// =========================================================
// TEST 7 — BORRAR: caso de éxito
// =========================================================

func TestOrdenService_Borrar_Exito(t *testing.T) {
	repo := new(mockOrdenRepo)
	repo.On("BorrarOrden", 1).Return(true)
	svc := service.NewOrdenCorrectivaService(repo)

	err := svc.Borrar(1)

	require.NoError(t, err)
}
