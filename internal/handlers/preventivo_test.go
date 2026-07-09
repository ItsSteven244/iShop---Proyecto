package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ItsSteven244/iShop---Proyecto/internal/handlers"
	"github.com/ItsSteven244/iShop---Proyecto/internal/middleware"
	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

type fakeMantenimientoRepo struct {
	mantenimientos []models.MantenimientoPreventivo
	nextID         int
}

func newFakeMantenimientoRepo() *fakeMantenimientoRepo {
	return &fakeMantenimientoRepo{nextID: 1}
}
func (f *fakeMantenimientoRepo) ListarMantenimientos() []models.MantenimientoPreventivo {
	return f.mantenimientos
}
func (f *fakeMantenimientoRepo) BuscarMantenimientoPorID(id int) (models.MantenimientoPreventivo, bool) {
	for _, m := range f.mantenimientos {
		if m.ID == id {
			return m, true
		}
	}
	return models.MantenimientoPreventivo{}, false
}
func (f *fakeMantenimientoRepo) CrearMantenimiento(mant models.MantenimientoPreventivo) models.MantenimientoPreventivo {
	mant.ID = f.nextID
	f.nextID++
	f.mantenimientos = append(f.mantenimientos, mant)
	return mant
}
func (f *fakeMantenimientoRepo) ActualizarMantenimiento(id int, datos models.MantenimientoPreventivo) (models.MantenimientoPreventivo, bool) {
	return models.MantenimientoPreventivo{}, false
}
func (f *fakeMantenimientoRepo) BorrarMantenimiento(id int) bool { return false }

func setupRouterMantenimientos(s *handlers.Server, authSvc *service.AuthService) http.Handler {
	r := chi.NewRouter()
	r.Use(chimw.Recoverer)
	r.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))
			r.Get("/mantenimientos", s.ListarMantenimientos)
			r.Post("/mantenimientos", s.CrearMantenimiento)
		})
	})
	return r
}

func TestHandler_CrearMantenimiento_Exitoso(t *testing.T) {
	fakeRepo := newFakeMantenimientoRepo()
	mantSvc := service.NewMantenimientoPreventivoService(fakeRepo)
	usuarioRepo := &fakeUsuarioRepo{}
	authSvc := service.NewAuthService(usuarioRepo)

	servidor := handlers.NewServer(nil, nil, nil, authSvc, nil, nil, nil, mantSvc, nil, nil)
	router := setupRouterMantenimientos(servidor, authSvc)

	usuario, _ := authSvc.Registrar("preventivo@test.com", "123456", "tecnico")
	token, _ := authSvc.GenerarToken(usuario)

	cuerpo, _ := json.Marshal(models.MantenimientoPreventivo{
		Equipo:            "Computadora",
		FechaProgramada:   "2026-07-01",
		TipoMantenimiento: "Preventivo",
		Estado:            "Pendiente",
		TecnicoID:         1,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/mantenimientos", bytes.NewReader(cuerpo))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
}

func TestHandler_ListarMantenimientos_SinToken_401(t *testing.T) {
	fakeRepo := newFakeMantenimientoRepo()
	mantSvc := service.NewMantenimientoPreventivoService(fakeRepo)
	usuarioRepo := &fakeUsuarioRepo{}
	authSvc := service.NewAuthService(usuarioRepo)

	servidor := handlers.NewServer(nil, nil, nil, authSvc, nil, nil, nil, mantSvc, nil, nil)
	router := setupRouterMantenimientos(servidor, authSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/mantenimientos", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
