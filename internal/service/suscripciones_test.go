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

// =========================================================
// TEST 1 — CREAR: precio inválido, no debe llegar al repo
// =========================================================

func TestServicioDigitalService_Crear_PrecioInvalido(t *testing.T) {
	repo := new(mockServicioRepo)
	repo.On("CrearServicio", mock.Anything).Return(models.ServicioDigital{})
	svc := service.NewServicioDigitalService(repo)

	_, err := svc.Crear(models.ServicioDigital{
		Nombre: "Netflix",
		Precio: 0,
	})

	require.ErrorIs(t, err, service.ErrPrecioInvalido)
	repo.AssertNotCalled(t, "CrearServicio")
}

// =========================================================
// TEST 2 — CREAR: nombre vacío, no debe llegar al repo
// =========================================================

func TestServicioDigitalService_Crear_NombreVacio(t *testing.T) {
	repo := new(mockServicioRepo)
	repo.On("CrearServicio", mock.Anything).Return(models.ServicioDigital{})
	svc := service.NewServicioDigitalService(repo)

	_, err := svc.Crear(models.ServicioDigital{
		Nombre: "",
		Precio: 9.99,
	})

	require.ErrorIs(t, err, service.ErrNombreServicioVacio)
	repo.AssertNotCalled(t, "CrearServicio")
}

// =========================================================
// TEST 3 — CREAR: caso de éxito
// =========================================================

func TestServicioDigitalService_Crear_Exito(t *testing.T) {
	repo := new(mockServicioRepo)
	entrada := models.ServicioDigital{Nombre: "Netflix", Categoria: "Streaming", Precio: 9.99, DuracionDias: 30, CantidadPerfiles: 4}
	esperado := models.ServicioDigital{ID: 1, Nombre: "Netflix", Categoria: "Streaming", Precio: 9.99, DuracionDias: 30, CantidadPerfiles: 4}

	repo.On("CrearServicio", entrada).Return(esperado)
	svc := service.NewServicioDigitalService(repo)

	resultado, err := svc.Crear(entrada)

	require.NoError(t, err)
	require.Equal(t, esperado, resultado)
	repo.AssertExpectations(t)
}

// =========================================================
// TEST 4 y 5 — OBTENER: encontrado y no encontrado
// =========================================================

func TestServicioDigitalService_Obtener_Encontrado(t *testing.T) {
	repo := new(mockServicioRepo)
	esperado := models.ServicioDigital{ID: 1, Nombre: "Netflix"}
	repo.On("BuscarServicioPorID", 1).Return(esperado, true)
	svc := service.NewServicioDigitalService(repo)

	resultado, err := svc.Obtener(1)

	require.NoError(t, err)
	require.Equal(t, esperado, resultado)
}

func TestServicioDigitalService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(mockServicioRepo)
	repo.On("BuscarServicioPorID", 999).Return(models.ServicioDigital{}, false)
	svc := service.NewServicioDigitalService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

// =========================================================
// TEST 6 — ACTUALIZAR: caso de éxito
// =========================================================

func TestServicioDigitalService_Actualizar_Exito(t *testing.T) {
	repo := new(mockServicioRepo)
	datos := models.ServicioDigital{Nombre: "Netflix Premium", Categoria: "Streaming", Precio: 12.99, DuracionDias: 30, CantidadPerfiles: 4}
	actualizado := models.ServicioDigital{ID: 1, Nombre: "Netflix Premium", Categoria: "Streaming", Precio: 12.99, DuracionDias: 30, CantidadPerfiles: 4}

	repo.On("ActualizarServicio", 1, datos).Return(actualizado, true)
	svc := service.NewServicioDigitalService(repo)

	resultado, err := svc.Actualizar(1, datos)

	require.NoError(t, err)
	require.Equal(t, actualizado, resultado)
}

// =========================================================
// TEST 7 — ACTUALIZAR: no encontrado
// =========================================================

func TestServicioDigitalService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(mockServicioRepo)
	datos := models.ServicioDigital{Nombre: "Netflix", Precio: 9.99}
	repo.On("ActualizarServicio", 999, datos).Return(models.ServicioDigital{}, false)
	svc := service.NewServicioDigitalService(repo)

	_, err := svc.Actualizar(999, datos)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

// =========================================================
// TEST 8 — BORRAR: caso de éxito
// =========================================================

func TestServicioDigitalService_Borrar_Exito(t *testing.T) {
	repo := new(mockServicioRepo)
	repo.On("BorrarServicio", 1).Return(true)
	svc := service.NewServicioDigitalService(repo)

	err := svc.Borrar(1)

	require.NoError(t, err)
}

// =========================================================
// TEST 9 — BORRAR: no encontrado
// =========================================================

func TestServicioDigitalService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(mockServicioRepo)
	repo.On("BorrarServicio", 999).Return(false)
	svc := service.NewServicioDigitalService(repo)

	err := svc.Borrar(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}
