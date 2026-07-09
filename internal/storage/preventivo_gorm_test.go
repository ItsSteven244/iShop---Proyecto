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

// =========================================================
// TEST 1 — REPOSITORIO GORM CON :memory:
// =========================================================

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

// =========================================================
// TEST 2 — REPOSITORIO GORM: BUSCAR MANTENIMIENTO INEXISTENTE
// =========================================================

func TestPreventivoGORM_BuscarMantenimientoPorID_NoExiste(t *testing.T) {
	// Preparar — base vacía, sin ningún mantenimiento creado
	db := abrirDBMemoriaPreventivo(t)
	repo := storage.NuevoPreventivoGORM(db)

	// Ejecutar — buscamos un ID que nunca se creó
	_, ok := repo.BuscarMantenimientoPorID(999)

	// Verificar — debe indicar que no se encontró
	require.False(t, ok)
}
