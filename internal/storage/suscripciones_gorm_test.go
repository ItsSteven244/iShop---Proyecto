package storage_test

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

func abrirDBMemoriaSuscripciones(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(&models.ServicioDigital{}, &models.SuscripcionCliente{}, &models.AccesoDigital{})
	require.NoError(t, err)
	return db
}

// =========================================================
// SERVICIO DIGITAL
// =========================================================

func TestSuscripcionesGORM_CrearServicio_Refleja(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	servicio := models.ServicioDigital{
		Nombre:           "Netflix",
		Categoria:        "Streaming",
		Precio:           9.99,
		DuracionDias:     30,
		CantidadPerfiles: 4,
	}

	creado := repo.CrearServicio(servicio)

	if creado.ID == 0 {
		t.Fatalf("la base no asignó ID al servicio")
	}

	encontrado, ok := repo.BuscarServicioPorID(creado.ID)
	require.True(t, ok)
	require.Equal(t, "Netflix", encontrado.Nombre)

	lista := repo.ListarServicios()
	require.Len(t, lista, 1)
}

func TestSuscripcionesGORM_BuscarServicioPorID_NoExiste(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	_, ok := repo.BuscarServicioPorID(999)

	require.False(t, ok)
}

func TestSuscripcionesGORM_ActualizarServicio(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	creado := repo.CrearServicio(models.ServicioDigital{Nombre: "Netflix", Precio: 9.99})

	actualizado, ok := repo.ActualizarServicio(creado.ID, models.ServicioDigital{Nombre: "Netflix Premium", Precio: 12.99})

	require.True(t, ok)
	require.Equal(t, "Netflix Premium", actualizado.Nombre)
}

func TestSuscripcionesGORM_BorrarServicio(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	creado := repo.CrearServicio(models.ServicioDigital{Nombre: "Netflix", Precio: 9.99})

	ok := repo.BorrarServicio(creado.ID)
	require.True(t, ok)

	_, existe := repo.BuscarServicioPorID(creado.ID)
	require.False(t, existe)
}

// =========================================================
// SUSCRIPCION CLIENTE
// =========================================================

func TestSuscripcionesGORM_CrearSuscripcion_Refleja(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	suscripcion := models.SuscripcionCliente{
		ClienteID:         1,
		ServicioDigitalID: 1,
		FechaInicio:       "2026-07-01",
		FechaFin:          "2026-08-01",
		Estado:            "Activa",
		TecnicoID:         1,
	}

	creada := repo.CrearSuscripcion(suscripcion)

	if creada.ID == 0 {
		t.Fatalf("la base no asignó ID a la suscripción")
	}

	encontrada, ok := repo.BuscarSuscripcionPorID(creada.ID)
	require.True(t, ok)
	require.Equal(t, "Activa", encontrada.Estado)

	lista := repo.ListarSuscripciones()
	require.Len(t, lista, 1)
}

func TestSuscripcionesGORM_BuscarSuscripcionPorID_NoExiste(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	_, ok := repo.BuscarSuscripcionPorID(999)

	require.False(t, ok)
}

func TestSuscripcionesGORM_ActualizarSuscripcion(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	creada := repo.CrearSuscripcion(models.SuscripcionCliente{
		ClienteID: 1, ServicioDigitalID: 1,
		FechaInicio: "2026-07-01", FechaFin: "2026-08-01",
		Estado: "Activa", TecnicoID: 1,
	})

	actualizada, ok := repo.ActualizarSuscripcion(creada.ID, models.SuscripcionCliente{
		ClienteID: 1, ServicioDigitalID: 1,
		FechaInicio: "2026-07-01", FechaFin: "2026-08-01",
		Estado: "Renovada", TecnicoID: 1,
	})

	require.True(t, ok)
	require.Equal(t, "Renovada", actualizada.Estado)
}

func TestSuscripcionesGORM_BorrarSuscripcion(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	creada := repo.CrearSuscripcion(models.SuscripcionCliente{
		ClienteID: 1, ServicioDigitalID: 1,
		FechaInicio: "2026-07-01", FechaFin: "2026-08-01",
		Estado: "Activa", TecnicoID: 1,
	})

	ok := repo.BorrarSuscripcion(creada.ID)
	require.True(t, ok)

	_, existe := repo.BuscarSuscripcionPorID(creada.ID)
	require.False(t, existe)
}

// =========================================================
// ACCESO DIGITAL
// =========================================================

func TestSuscripcionesGORM_CrearAcceso_Refleja(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	acceso := models.AccesoDigital{
		CorreoAcceso:         "cliente@correo.com",
		Perfil:               "Perfil 1",
		Estado:               "Activo",
		SuscripcionClienteID: 1,
	}

	creado := repo.CrearAcceso(acceso)

	if creado.ID == 0 {
		t.Fatalf("la base no asignó ID al acceso")
	}

	encontrado, ok := repo.BuscarAccesoPorID(creado.ID)
	require.True(t, ok)
	require.Equal(t, "cliente@correo.com", encontrado.CorreoAcceso)

	lista := repo.ListarAccesos()
	require.Len(t, lista, 1)
}

func TestSuscripcionesGORM_BuscarAccesoPorID_NoExiste(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	_, ok := repo.BuscarAccesoPorID(999)

	require.False(t, ok)
}

func TestSuscripcionesGORM_ActualizarAcceso(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	creado := repo.CrearAcceso(models.AccesoDigital{
		CorreoAcceso: "cliente@correo.com", Perfil: "Perfil 1",
		Estado: "Activo", SuscripcionClienteID: 1,
	})

	actualizado, ok := repo.ActualizarAcceso(creado.ID, models.AccesoDigital{
		CorreoAcceso: "cliente@correo.com", Perfil: "Perfil 1",
		Estado: "Suspendido", SuscripcionClienteID: 1,
	})

	require.True(t, ok)
	require.Equal(t, "Suspendido", actualizado.Estado)
}

func TestSuscripcionesGORM_BorrarAcceso(t *testing.T) {
	db := abrirDBMemoriaSuscripciones(t)
	repo := storage.NuevoSuscripcionesGORM(db)

	creado := repo.CrearAcceso(models.AccesoDigital{
		CorreoAcceso: "cliente@correo.com", Perfil: "Perfil 1",
		Estado: "Activo", SuscripcionClienteID: 1,
	})

	ok := repo.BorrarAcceso(creado.ID)
	require.True(t, ok)

	_, existe := repo.BuscarAccesoPorID(creado.ID)
	require.False(t, existe)
}
