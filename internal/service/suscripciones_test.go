package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

// =========================================================
// MOCK DEL REPOSITORIO — ServicioDigital
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

// =========================================================
// TEST 10 — LISTAR (estaba en 0%, esta es la ganancia rápida)
// =========================================================

func TestServicioDigitalService_Listar(t *testing.T) {
	repo := new(mockServicioRepo)
	esperado := []models.ServicioDigital{
		{ID: 1, Nombre: "Netflix"},
		{ID: 2, Nombre: "Spotify"},
	}
	repo.On("ListarServicios").Return(esperado)
	svc := service.NewServicioDigitalService(repo)

	resultado := svc.Listar()

	require.Equal(t, esperado, resultado)
	repo.AssertExpectations(t)
}

// =========================================================
// MOCK DEL REPOSITORIO — SuscripcionCliente
// =========================================================

type mockSuscripcionRepo struct {
	mock.Mock
}

func (m *mockSuscripcionRepo) ListarSuscripciones() []models.SuscripcionCliente {
	args := m.Called()
	return args.Get(0).([]models.SuscripcionCliente)
}
func (m *mockSuscripcionRepo) BuscarSuscripcionPorID(id int) (models.SuscripcionCliente, bool) {
	args := m.Called(id)
	return args.Get(0).(models.SuscripcionCliente), args.Bool(1)
}
func (m *mockSuscripcionRepo) CrearSuscripcion(s models.SuscripcionCliente) models.SuscripcionCliente {
	args := m.Called(s)
	return args.Get(0).(models.SuscripcionCliente)
}
func (m *mockSuscripcionRepo) ActualizarSuscripcion(id int, datos models.SuscripcionCliente) (models.SuscripcionCliente, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.SuscripcionCliente), args.Bool(1)
}
func (m *mockSuscripcionRepo) BorrarSuscripcion(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// =========================================================
// TEST 11 a 15 — SuscripcionCliente: validaciones de Crear
// (validarSuscripcion revisa 6 campos en cadena: probamos los
// primeros dos, que son los que más te van a preguntar en la
// defensa por ser los "obligatorios" de relación)
// =========================================================

func TestSuscripcionClienteService_Crear_ClienteIDVacio(t *testing.T) {
	repo := new(mockSuscripcionRepo)
	repo.On("CrearSuscripcion", mock.Anything).Return(models.SuscripcionCliente{})
	svc := service.NewSuscripcionClienteService(repo)

	_, err := svc.Crear(models.SuscripcionCliente{
		ClienteID:         0,
		ServicioDigitalID: 1,
	})

	require.ErrorIs(t, err, service.ErrClienteIDRequerido)
	repo.AssertNotCalled(t, "CrearSuscripcion")
}

func TestSuscripcionClienteService_Crear_ServicioDigitalIDVacio(t *testing.T) {
	repo := new(mockSuscripcionRepo)
	repo.On("CrearSuscripcion", mock.Anything).Return(models.SuscripcionCliente{})
	svc := service.NewSuscripcionClienteService(repo)

	_, err := svc.Crear(models.SuscripcionCliente{
		ClienteID:         1,
		ServicioDigitalID: 0,
	})

	require.ErrorIs(t, err, service.ErrServicioDigitalIDRequerido)
	repo.AssertNotCalled(t, "CrearSuscripcion")
}

func TestSuscripcionClienteService_Crear_Exito(t *testing.T) {
	repo := new(mockSuscripcionRepo)
	entrada := models.SuscripcionCliente{
		ClienteID:         1,
		ServicioDigitalID: 1,
		FechaInicio:       "2026-07-01",
		FechaFin:          "2026-08-01",
		Estado:            "Activa",
		TecnicoID:         1,
	}
	esperado := entrada
	esperado.ID = 1

	repo.On("CrearSuscripcion", entrada).Return(esperado)
	svc := service.NewSuscripcionClienteService(repo)

	resultado, err := svc.Crear(entrada)

	require.NoError(t, err)
	require.Equal(t, esperado, resultado)
	repo.AssertExpectations(t)
}

func TestSuscripcionClienteService_Obtener_Encontrado(t *testing.T) {
	repo := new(mockSuscripcionRepo)
	esperado := models.SuscripcionCliente{ID: 1, ClienteID: 1}
	repo.On("BuscarSuscripcionPorID", 1).Return(esperado, true)
	svc := service.NewSuscripcionClienteService(repo)

	resultado, err := svc.Obtener(1)

	require.NoError(t, err)
	require.Equal(t, esperado, resultado)
}

func TestSuscripcionClienteService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(mockSuscripcionRepo)
	repo.On("BuscarSuscripcionPorID", 999).Return(models.SuscripcionCliente{}, false)
	svc := service.NewSuscripcionClienteService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

func TestSuscripcionClienteService_Actualizar_Exito(t *testing.T) {
	repo := new(mockSuscripcionRepo)
	datos := models.SuscripcionCliente{
		ClienteID: 1, ServicioDigitalID: 1,
		FechaInicio: "2026-07-01", FechaFin: "2026-08-01",
		Estado: "Renovada", TecnicoID: 1,
	}
	actualizada := datos
	actualizada.ID = 1

	repo.On("ActualizarSuscripcion", 1, datos).Return(actualizada, true)
	svc := service.NewSuscripcionClienteService(repo)

	resultado, err := svc.Actualizar(1, datos)

	require.NoError(t, err)
	require.Equal(t, actualizada, resultado)
}

func TestSuscripcionClienteService_Borrar_Exito(t *testing.T) {
	repo := new(mockSuscripcionRepo)
	repo.On("BorrarSuscripcion", 1).Return(true)
	svc := service.NewSuscripcionClienteService(repo)

	err := svc.Borrar(1)

	require.NoError(t, err)
}

func TestSuscripcionClienteService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(mockSuscripcionRepo)
	repo.On("BorrarSuscripcion", 999).Return(false)
	svc := service.NewSuscripcionClienteService(repo)

	err := svc.Borrar(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

// =========================================================
// MOCK DEL REPOSITORIO — AccesoDigital
// =========================================================

type mockAccesoRepo struct {
	mock.Mock
}

func (m *mockAccesoRepo) ListarAccesos() []models.AccesoDigital {
	args := m.Called()
	return args.Get(0).([]models.AccesoDigital)
}
func (m *mockAccesoRepo) BuscarAccesoPorID(id int) (models.AccesoDigital, bool) {
	args := m.Called(id)
	return args.Get(0).(models.AccesoDigital), args.Bool(1)
}
func (m *mockAccesoRepo) CrearAcceso(a models.AccesoDigital) models.AccesoDigital {
	args := m.Called(a)
	return args.Get(0).(models.AccesoDigital)
}
func (m *mockAccesoRepo) ActualizarAcceso(id int, datos models.AccesoDigital) (models.AccesoDigital, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.AccesoDigital), args.Bool(1)
}
func (m *mockAccesoRepo) BorrarAcceso(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func TestAccesoDigitalService_Crear_CorreoAccesoVacio(t *testing.T) {
	repo := new(mockAccesoRepo)
	repo.On("CrearAcceso", mock.Anything).Return(models.AccesoDigital{})
	svc := service.NewAccesoDigitalService(repo)

	_, err := svc.Crear(models.AccesoDigital{
		CorreoAcceso: "",
		Perfil:       "Perfil 1",
	})

	require.ErrorIs(t, err, service.ErrCorreoAccesoVacio)
	repo.AssertNotCalled(t, "CrearAcceso")
}

func TestAccesoDigitalService_Crear_PerfilVacio(t *testing.T) {
	repo := new(mockAccesoRepo)
	repo.On("CrearAcceso", mock.Anything).Return(models.AccesoDigital{})
	svc := service.NewAccesoDigitalService(repo)

	_, err := svc.Crear(models.AccesoDigital{
		CorreoAcceso: "cliente@correo.com",
		Perfil:       "",
	})

	require.ErrorIs(t, err, service.ErrPerfilVacio)
	repo.AssertNotCalled(t, "CrearAcceso")
}

func TestAccesoDigitalService_Crear_Exito(t *testing.T) {
	repo := new(mockAccesoRepo)
	entrada := models.AccesoDigital{
		CorreoAcceso:         "cliente@correo.com",
		Perfil:               "Perfil 1",
		Estado:               "Activo",
		SuscripcionClienteID: 1,
	}
	esperado := entrada
	esperado.ID = 1

	repo.On("CrearAcceso", entrada).Return(esperado)
	svc := service.NewAccesoDigitalService(repo)

	resultado, err := svc.Crear(entrada)

	require.NoError(t, err)
	require.Equal(t, esperado, resultado)
	repo.AssertExpectations(t)
}

func TestAccesoDigitalService_Obtener_Encontrado(t *testing.T) {
	repo := new(mockAccesoRepo)
	esperado := models.AccesoDigital{ID: 1, CorreoAcceso: "cliente@correo.com"}
	repo.On("BuscarAccesoPorID", 1).Return(esperado, true)
	svc := service.NewAccesoDigitalService(repo)

	resultado, err := svc.Obtener(1)

	require.NoError(t, err)
	require.Equal(t, esperado, resultado)
}

func TestAccesoDigitalService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(mockAccesoRepo)
	repo.On("BuscarAccesoPorID", 999).Return(models.AccesoDigital{}, false)
	svc := service.NewAccesoDigitalService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}

func TestAccesoDigitalService_Actualizar_Exito(t *testing.T) {
	repo := new(mockAccesoRepo)
	datos := models.AccesoDigital{
		CorreoAcceso: "cliente@correo.com", Perfil: "Perfil 1",
		Estado: "Suspendido", SuscripcionClienteID: 1,
	}
	actualizado := datos
	actualizado.ID = 1

	repo.On("ActualizarAcceso", 1, datos).Return(actualizado, true)
	svc := service.NewAccesoDigitalService(repo)

	resultado, err := svc.Actualizar(1, datos)

	require.NoError(t, err)
	require.Equal(t, actualizado, resultado)
}

func TestAccesoDigitalService_Borrar_Exito(t *testing.T) {
	repo := new(mockAccesoRepo)
	repo.On("BorrarAcceso", 1).Return(true)
	svc := service.NewAccesoDigitalService(repo)

	err := svc.Borrar(1)

	require.NoError(t, err)
}

func TestAccesoDigitalService_Borrar_NoEncontrado(t *testing.T) {
	repo := new(mockAccesoRepo)
	repo.On("BorrarAcceso", 999).Return(false)
	svc := service.NewAccesoDigitalService(repo)

	err := svc.Borrar(999)

	require.ErrorIs(t, err, service.ErrNoEncontrado)
}
