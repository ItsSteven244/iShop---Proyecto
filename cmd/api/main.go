package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"github.com/ItsSteven244/iShop---Proyecto/internal/handlers"
	"github.com/ItsSteven244/iShop---Proyecto/internal/middleware"
	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

func main() {
	// 1. Abrir la base de datos SQLite con GORM y migrar los modelos.
	db, err := gorm.Open(sqlite.Open("ishop.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := db.AutoMigrate(
		&models.OrdenCorrectiva{},
		&models.ProcesoReparacion{},
		&models.EvidenciaDanio{},
		&models.Usuario{},
		// compañeros agregan sus modelos aquí
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Crear los repositorios GORM.
	correctivoRepo := storage.NuevoCorrectivoGORM(db)
	usuarioRepo := storage.NewUsuarioGORM(db)
	// compañeros agregan sus repositorios aquí

	// 3. Crear los servicios.
	ordenService := service.NewOrdenCorrectivaService(correctivoRepo)
	procesoService := service.NewProcesoReparacionService(correctivoRepo)
	evidenciaService := service.NewEvidenciaDanioService(correctivoRepo)
	authService := service.NewAuthService(usuarioRepo)
	// compañeros agregan sus servicios aquí

	// 4. Crear el servidor con inyección de dependencias.
	servidor := handlers.NewServer(ordenService, procesoService, evidenciaService, authService)

	// 5. Configurar el router principal.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 6. Registrar las rutas.
	r.Route("/api/v1", func(r chi.Router) {
		// Rutas públicas
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		// Rutas protegidas
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			// Módulo Correctivo
			r.Get("/correctivos/ordenes", servidor.ListarOrdenes)
			r.Post("/correctivos/ordenes", servidor.CrearOrden)
			r.Get("/correctivos/ordenes/{id}", servidor.ObtenerOrden)
			r.Put("/correctivos/ordenes/{id}", servidor.ActualizarOrden)
			r.Patch("/correctivos/ordenes/{id}", servidor.ActualizarEstadoOrden)
			r.Delete("/correctivos/ordenes/{id}", servidor.BorrarOrden)

			r.Get("/correctivos/procesos", servidor.ListarProcesos)
			r.Post("/correctivos/procesos", servidor.CrearProceso)
			r.Get("/correctivos/procesos/{id}", servidor.ObtenerProceso)
			r.Put("/correctivos/procesos/{id}", servidor.ActualizarProceso)
			r.Delete("/correctivos/procesos/{id}", servidor.BorrarProceso)

			r.Get("/correctivos/evidencias", servidor.ListarEvidencias)
			r.Post("/correctivos/evidencias", servidor.CrearEvidencia)
			r.Get("/correctivos/evidencias/{id}", servidor.ObtenerEvidencia)
			r.Put("/correctivos/evidencias/{id}", servidor.ActualizarEvidencia)
			r.Delete("/correctivos/evidencias/{id}", servidor.BorrarEvidencia)

			// Módulo Preventivo - compañero agrega sus rutas aquí

			// Módulo Suscripciones - compañero agrega sus rutas aquí
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
