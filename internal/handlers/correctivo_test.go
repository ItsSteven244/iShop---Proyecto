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

// fakeOrdenRepo es un fake en memoria para los tests de handler.
type fakeOrdenRepo struct {
	ordenes []models.OrdenCorrectiva
	nextID  int
}

func newFakeOrdenRepo() *fakeOrdenRepo {
	return &fakeOrdenRepo{nextID: 1}
}
func (f *fakeOrdenRepo) ListarOrdenes() []models.OrdenCorrectiva { return f.ordenes }
func (f *fakeOrdenRepo) BuscarOrdenPorID(id int) (models.OrdenCorrectiva, bool) {
	for _, o := range f.ordenes {
		if o.ID == id {
			return o, true
		}
	}
	return models.OrdenCorrectiva{}, false
}
func (f *fakeOrdenRepo) CrearOrden(o models.OrdenCorrectiva) models.OrdenCorrectiva {
	o.ID = f.nextID
	f.nextID++
	f.ordenes = append(f.ordenes, o)
	return o
}
func (f *fakeOrdenRepo) ActualizarOrden(id int, datos models.OrdenCorrectiva) (models.OrdenCorrectiva, bool) {
	return models.OrdenCorrectiva{}, false
}
func (f *fakeOrdenRepo) ActualizarOrdenParcial(id int, estado string, diagnostico string) (models.OrdenCorrectiva, bool) {
	return models.OrdenCorrectiva{}, false
}
func (f *fakeOrdenRepo) BorrarOrden(id int) bool { return false }

// fakeUsuarioRepo es un fake del repositorio de usuarios para el authService.
type fakeUsuarioRepo struct {
	usuarios []models.Usuario
}

func (f *fakeUsuarioRepo) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.ID = len(f.usuarios) + 1
	f.usuarios = append(f.usuarios, u)
	return u, nil
}
func (f *fakeUsuarioRepo) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	for _, u := range f.usuarios {
		if u.Email == email {
			return u, true
		}
	}
	return models.Usuario{}, false
}

// setupRouter arma el router con middleware de auth para los tests.
func setupRouter(s *handlers.Server, authSvc *service.AuthService) http.Handler {
	r := chi.NewRouter()
	r.Use(chimw.Recoverer)
	r.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))
			r.Get("/correctivos/ordenes", s.ListarOrdenes)
			r.Post("/correctivos/ordenes", s.CrearOrden)
		})
	})
	return r
}

// =========================================================
// TEST 2 — HANDLER CON HTTPTEST (crear orden con token)
// =========================================================

func TestHandler_CrearOrden_Exitoso(t *testing.T) {
	// Preparar
	fakeRepo := newFakeOrdenRepo()
	ordenSvc := service.NewOrdenCorrectivaService(fakeRepo)
	usuarioRepo := &fakeUsuarioRepo{}
	authSvc := service.NewAuthService(usuarioRepo)

	servidor := handlers.NewServer(ordenSvc, nil, nil, authSvc)
	router := setupRouter(servidor, authSvc)

	// Generar token válido
	usuario, _ := authSvc.Registrar("test@test.com", "123456")
	token, _ := authSvc.GenerarToken(usuario)

	cuerpo, _ := json.Marshal(models.OrdenCorrectiva{
		Codigo:            "ORD-001",
		ProblemaReportado: "Pantalla rota",
	})

	// Ejecutar
	req := httptest.NewRequest(http.MethodPost, "/api/v1/correctivos/ordenes", bytes.NewReader(cuerpo))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Verificar
	require.Equal(t, http.StatusCreated, rec.Code)
}

// =========================================================
// TEST 3 — HANDLER 401 SIN TOKEN
// =========================================================

func TestHandler_ListarOrdenes_SinToken_401(t *testing.T) {
	// Preparar
	fakeRepo := newFakeOrdenRepo()
	ordenSvc := service.NewOrdenCorrectivaService(fakeRepo)
	usuarioRepo := &fakeUsuarioRepo{}
	authSvc := service.NewAuthService(usuarioRepo)

	servidor := handlers.NewServer(ordenSvc, nil, nil, authSvc)
	router := setupRouter(servidor, authSvc)

	// Ejecutar — sin token
	req := httptest.NewRequest(http.MethodGet, "/api/v1/correctivos/ordenes", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Verificar
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
