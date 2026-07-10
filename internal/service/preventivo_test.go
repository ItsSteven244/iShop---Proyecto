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

// =========================================================
// TEST 1 — CREAR: equipo vacío, no debe llegar al repo
// =========================================================

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

// =========================================================
// TEST 2 — CREAR: fecha programada vacía, no debe llegar al repo
// =========================================================

func TestMantenimientoPreventivoService_Crear_FechaProgramadaVacia(t *testing.T) {
	repo := new(mockMantenimientoRepo)
	repo.On("CrearMantenimiento", mock.Anything).Return(models.MantenimientoPreventivo{})
	svc := service.NewMantenimientoPreventivoService(repo)

	_, err := svc.Crear(models.MantenimientoPreventivo{
		Equipo:            "Computadora",
		FechaProgramada:   "",
		TipoMantenimiento: "Preventivo",
	})

	require.ErrorIs(t, err, service.ErrFechaProgramadaVacia)
	repo.AssertNotCalled(t, "CrearMantenimiento")
}

// =========================================================
// TEST 3 — CREAR: caso de éxito.
// =========================================================

func TestMantenimientoPreventivoService_Crear_Exito(t *testing.T) {
	repo := new(mockMantenimientoRepo)
	entrada := models.MantenimientoPreventivo{Equipo: "Computadora", FechaProgramada: "2026-07-01", TipoMantenimiento: "Preventivo", Estado: "Pendiente", TecnicoID: 1}
	esperado := models.MantenimientoPreventivo{ID: 1, Equipo: "Computadora", FechaProgramada: "2026-07-01", TipoMantenimiento: "Preventivo", Estado: "Pendiente", TecnicoID: 1}

	repo.On("CrearMantenimiento", entrada).Return(esperado)
	svc := service.NewMantenimientoPreventivoService(repo)

	resultado, err := svc.Crear(entrada)

	require.NoError(t, err)
	require.Equal(t, esperado, resultado)
	repo.AssertExpectations(t)
}

// =========================================================
// TEST 4 y 5 — OBTENER: encontrado y no encontrado
// =========================================================

func TestMantenimientoPreventivoService_Obtener_Encontrado(t *testing.T) {
	repo := new(mockMantenimientoRepo)
	esperado := models.MantenimientoPreventivo{ID: 1, Equipo: "Computadora"}
	repo.On("BuscarMantenimientoPorID", 1).Return(esperado, true)
	svc := service.NewMantenimientoPreventivoService(repo)

	resultado, err := svc.Obtener(1)

	require.NoError(t, err)
	require.Equal(t, esperado, resultado)
}

func TestMantenimientoPreventivoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(mockMantenimientoRepo)
	repo.On("BuscarMantenimientoPorID", 999).Return(models.MantenimientoPreventivo{}, false)
	svc := service.NewMantenimientoPreventivoService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

// =========================================================
// TEST 6 — ACTUALIZAR: caso de éxito
// =========================================================

func TestMantenimientoPreventivoService_Actualizar_Exito(t *testing.T) {
	repo := new(mockMantenimientoRepo)
	datos := models.MantenimientoPreventivo{Equipo: "Computadora", FechaProgramada: "2026-07-02", TipoMantenimiento: "Preventivo", Estado: "En proceso", TecnicoID: 1}
	actualizado := models.MantenimientoPreventivo{ID: 1, Equipo: "Computadora", FechaProgramada: "2026-07-02", TipoMantenimiento: "Preventivo", Estado: "En proceso", TecnicoID: 1}

	repo.On("ActualizarMantenimiento", 1, datos).Return(actualizado, true)
	svc := service.NewMantenimientoPreventivoService(repo)

	resultado, err := svc.Actualizar(1, datos)

	require.NoError(t, err)
	require.Equal(t, actualizado, resultado)
}

// =========================================================
// TEST 7 — ACTUALIZAR: no encontrado
// =========================================================

func TestMantenimientoPreventivoService_Actualizar_NoEncontrado(t *testing.T) {
	repo := new(mockMantenimientoRepo)
	datos := models.MantenimientoPreventivo{Equipo: "Computadora", FechaProgramada: "2026-07-02"}
	repo.On("ActualizarMantenimiento", 999, datos).Return(models.MantenimientoPreventivo{}, false)
	svc := service.NewMantenimientoPreventivoService(repo)

	_, err := svc.Actualizar(999, datos)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

// =========================================================
// TEST 8 — BORRAR: caso de éxito
// =========================================================

func TestMantenimientoPreventivoService_Borrar_Exito(t *testing.T) {
	repo := new(mockMantenimientoRepo)
	repo.On("BorrarMantenimiento", 1).Return(true)
	svc := service.NewMantenimientoPreventivoService(repo)

	err := svc.Borrar(1)

	require.NoError(t, err)
}

// =========================================================
// TEST 9 — BORRAR: no encontrado
// =========================================================

func TestMantenimientoPreventivoService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(mockMantenimientoRepo)
	repo.On("BorrarMantenimiento", 999).Return(false)
	svc := service.NewMantenimientoPreventivoService(repo)

	err := svc.Borrar(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}
