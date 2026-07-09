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

// fakeServicioRepo es un fake en memoria para los tests de handler.
type fakeServicioRepo struct {
	servicios []models.ServicioDigital
	nextID    int
}

func newFakeServicioRepo() *fakeServicioRepo {
	return &fakeServicioRepo{nextID: 1}
}
func (f *fakeServicioRepo) ListarServicios() []models.ServicioDigital { return f.servicios }
func (f *fakeServicioRepo) BuscarServicioPorID(id int) (models.ServicioDigital, bool) {
	for _, s := range f.servicios {
		if s.ID == id {
			return s, true
		}
	}
	return models.ServicioDigital{}, false
}
func (f *fakeServicioRepo) CrearServicio(s models.ServicioDigital) models.ServicioDigital {
	s.ID = f.nextID
	f.nextID++
	f.servicios = append(f.servicios, s)
	return s
}
func (f *fakeServicioRepo) ActualizarServicio(id int, datos models.ServicioDigital) (models.ServicioDigital, bool) {
	return models.ServicioDigital{}, false
}
func (f *fakeServicioRepo) BorrarServicio(id int) bool { return false }

// setupRouterServicios arma el router con middleware de auth para los tests.
func setupRouterServicios(s *handlers.Server, authSvc *service.AuthService) http.Handler {
	r := chi.NewRouter()
	r.Use(chimw.Recoverer)
	r.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))
			r.Get("/servicios", s.ListarServicios)
			r.Post("/servicios", s.CrearServicio)
		})
	})
	return r
}

func TestHandler_CrearServicio_Exitoso(t *testing.T) {
	fakeRepo := newFakeServicioRepo()
	servicioSvc := service.NewServicioDigitalService(fakeRepo)
	usuarioRepo := &fakeUsuarioRepo{}
	authSvc := service.NewAuthService(usuarioRepo)

	servidor := handlers.NewServer(nil, nil, nil, authSvc, servicioSvc, nil, nil, nil, nil, nil)
	router := setupRouterServicios(servidor, authSvc)

	usuario, _ := authSvc.Registrar("servicios@test.com", "123456")
	token, _ := authSvc.GenerarToken(usuario)

	cuerpo, _ := json.Marshal(models.ServicioDigital{
		Nombre:           "Netflix",
		Categoria:        "Streaming",
		Precio:           9.99,
		DuracionDias:     30,
		CantidadPerfiles: 4,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/servicios", bytes.NewReader(cuerpo))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)
}

func TestHandler_ListarServicios_SinToken_401(t *testing.T) {
	fakeRepo := newFakeServicioRepo()
	servicioSvc := service.NewServicioDigitalService(fakeRepo)
	usuarioRepo := &fakeUsuarioRepo{}
	authSvc := service.NewAuthService(usuarioRepo)

	servidor := handlers.NewServer(nil, nil, nil, authSvc, servicioSvc, nil, nil, nil, nil, nil)
	router := setupRouterServicios(servidor, authSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/servicios", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
