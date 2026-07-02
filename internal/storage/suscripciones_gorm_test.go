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
