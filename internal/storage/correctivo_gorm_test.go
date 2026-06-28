package storage_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
	"github.com/glebarez/sqlite"
)

// abrirDBMemoria abre una BD SQLite en memoria y migra los modelos.
func abrirDBMemoria(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(&models.OrdenCorrectiva{}, &models.ProcesoReparacion{}, &models.EvidenciaDanio{})
	require.NoError(t, err)
	return db
}

// =========================================================
// TEST 4 — REPOSITORIO GORM CON :memory:
// =========================================================

// TestCorrectivoGORM_CrearOrden_Refleja prueba que crear una orden
// queda reflejada al buscarla y listarla.
func TestCorrectivoGORM_CrearOrden_Refleja(t *testing.T) {
	// Preparar
	db := abrirDBMemoria(t)
	repo := storage.NuevoCorrectivoGORM(db)

	orden := models.OrdenCorrectiva{
		Codigo:            "ORD-001",
		ProblemaReportado: "Pantalla rota",
	}

	// Ejecutar
	creada := repo.CrearOrden(orden)

	// Verificar — la BD asignó un ID
	if creada.ID == 0 {
		t.Fatalf("la base no asignó ID a la orden")
	}

	// Verificar — buscar por ID la refleja
	encontrada, ok := repo.BuscarOrdenPorID(creada.ID)
	require.True(t, ok)
	require.Equal(t, "ORD-001", encontrada.Codigo)

	// Verificar — listar también la refleja
	lista := repo.ListarOrdenes()
	require.Len(t, lista, 1)
}

//...
