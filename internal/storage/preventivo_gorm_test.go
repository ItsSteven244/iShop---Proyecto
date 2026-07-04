package storage_test

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

func abrirDBMemoriaPreventivo(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(&models.MantenimientoPreventivo{}, &models.TareaPreventiva{}, &models.InsumoPreventivo{})
	require.NoError(t, err)
	return db
}

func TestPreventivoGORM_CrearMantenimiento_Refleja(t *testing.T) {
	db := abrirDBMemoriaPreventivo(t)
	repo := storage.NuevoPreventivoGORM(db)

	mant := models.MantenimientoPreventivo{
		Equipo:            "Computadora",
		FechaProgramada:   "2026-07-01",
		TipoMantenimiento: "Preventivo",
		Estado:            "Pendiente",
		TecnicoID:         1,
	}

	creado := repo.CrearMantenimiento(mant)

	if creado.ID == 0 {
		t.Fatalf("la base no asignó ID al mantenimiento")
	}

	encontrado, ok := repo.BuscarMantenimientoPorID(creado.ID)
	require.True(t, ok)
	require.Equal(t, "Computadora", encontrado.Equipo)

	lista := repo.ListarMantenimientos()
	require.Len(t, lista, 1)
}
