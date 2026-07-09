package main

import (
	"log"
	"net/http"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ItsSteven244/iShop---Proyecto/internal/handlers"
	"github.com/ItsSteven244/iShop---Proyecto/internal/middleware"
	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/seed"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

func main() {
	// 1. Abrir la base de datos y migrar los modelos.
	//    DB_DRIVER=postgres -> usa PostgreSQL (Docker).
	//    Sin DB_DRIVER (o cualquier otro valor) -> usa SQLite (desarrollo local).
	var db *gorm.DB
	var err error

	switch os.Getenv("DB_DRIVER") {
	case "postgres":
		dsn := os.Getenv("DATABASE_URL")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default: // sqlite
		db, err = gorm.Open(sqlite.Open("ishop.db"), &gorm.Config{})
	}
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	if err := db.AutoMigrate(
		&models.OrdenCorrectiva{},
		&models.ProcesoReparacion{},
		&models.EvidenciaDanio{},
		&models.Usuario{},
		&models.ServicioDigital{},
		&models.SuscripcionCliente{},
		&models.AccesoDigital{},
		&models.MantenimientoPreventivo{},
		&models.TareaPreventiva{},
		&models.InsumoPreventivo{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 1.5. Sembrar datos de ejemplo si la base está vacía.
	if err := seed.Ejecutar(db); err != nil {
		log.Fatal("falló el seed: ", err)
	}

	// 2. Crear los repositorios GORM.
	correctivoRepo := storage.NuevoCorrectivoGORM(db)
	usuarioRepo := storage.NewUsuarioGORM(db)
	suscripcionesRepo := storage.NuevoSuscripcionesGORM(db)
	preventivoRepo := storage.NuevoPreventivoGORM(db)

	// 3. Crear los servicios.
	ordenService := service.NewOrdenCorrectivaService(correctivoRepo)
	procesoService := service.NewProcesoReparacionService(correctivoRepo)
	evidenciaService := service.NewEvidenciaDanioService(correctivoRepo)
	authService := service.NewAuthService(usuarioRepo)
	servicioDigitalService := service.NewServicioDigitalService(suscripcionesRepo)
	suscripcionClienteService := service.NewSuscripcionClienteService(suscripcionesRepo)
	accesoDigitalService := service.NewAccesoDigitalService(suscripcionesRepo)
	mantenimientoService := service.NewMantenimientoPreventivoService(preventivoRepo)
	tareaService := service.NewTareaPreventivaService(preventivoRepo)
	insumoService := service.NewInsumoPreventivoService(preventivoRepo)

	// 4. Crear el servidor con inyección de dependencias.
	servidor := handlers.NewServer(
		ordenService, procesoService, evidenciaService, authService,
		servicioDigitalService, suscripcionClienteService, accesoDigitalService,
		mantenimientoService, tareaService, insumoService,
	)

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

			// Solo admin puede borrar (aplica a todos los DELETE de abajo).
			soloAdmin := middleware.RequireRol("admin")

			// Módulo Correctivo
			r.Get("/correctivos/ordenes", servidor.ListarOrdenes)
			r.Post("/correctivos/ordenes", servidor.CrearOrden)
			r.Get("/correctivos/ordenes/{id}", servidor.ObtenerOrden)
			r.Put("/correctivos/ordenes/{id}", servidor.ActualizarOrden)
			r.Patch("/correctivos/ordenes/{id}", servidor.ActualizarEstadoOrden)
			r.With(soloAdmin).Delete("/correctivos/ordenes/{id}", servidor.BorrarOrden)

			r.Get("/correctivos/procesos", servidor.ListarProcesos)
			r.Post("/correctivos/procesos", servidor.CrearProceso)
			r.Get("/correctivos/procesos/{id}", servidor.ObtenerProceso)
			r.Put("/correctivos/procesos/{id}", servidor.ActualizarProceso)
			r.With(soloAdmin).Delete("/correctivos/procesos/{id}", servidor.BorrarProceso)

			r.Get("/correctivos/evidencias", servidor.ListarEvidencias)
			r.Post("/correctivos/evidencias", servidor.CrearEvidencia)
			r.Get("/correctivos/evidencias/{id}", servidor.ObtenerEvidencia)
			r.Put("/correctivos/evidencias/{id}", servidor.ActualizarEvidencia)
			r.With(soloAdmin).Delete("/correctivos/evidencias/{id}", servidor.BorrarEvidencia)

			// Módulo Preventivo
			r.Get("/mantenimientos", servidor.ListarMantenimientos)
			r.Post("/mantenimientos", servidor.CrearMantenimiento)
			r.Get("/mantenimientos/{id}", servidor.ObtenerMantenimiento)
			r.Put("/mantenimientos/{id}", servidor.ActualizarMantenimiento)
			r.With(soloAdmin).Delete("/mantenimientos/{id}", servidor.BorrarMantenimiento)

			r.Get("/tareas", servidor.ListarTareas)
			r.Post("/tareas", servidor.CrearTarea)
			r.Get("/tareas/{id}", servidor.ObtenerTarea)
			r.Put("/tareas/{id}", servidor.ActualizarTarea)
			r.With(soloAdmin).Delete("/tareas/{id}", servidor.BorrarTarea)

			r.Get("/insumos", servidor.ListarInsumos)
			r.Post("/insumos", servidor.CrearInsumo)
			r.Get("/insumos/{id}", servidor.ObtenerInsumo)
			r.Put("/insumos/{id}", servidor.ActualizarInsumo)
			r.With(soloAdmin).Delete("/insumos/{id}", servidor.BorrarInsumo)

			// Módulo Suscripciones
			r.Get("/servicios", servidor.ListarServicios)
			r.Post("/servicios", servidor.CrearServicio)
			r.Get("/servicios/{id}", servidor.ObtenerServicio)
			r.Put("/servicios/{id}", servidor.ActualizarServicio)
			r.With(soloAdmin).Delete("/servicios/{id}", servidor.BorrarServicio)

			r.Get("/suscripciones", servidor.ListarSuscripciones)
			r.Post("/suscripciones", servidor.CrearSuscripcion)
			r.Get("/suscripciones/{id}", servidor.ObtenerSuscripcion)
			r.Put("/suscripciones/{id}", servidor.ActualizarSuscripcion)
			r.With(soloAdmin).Delete("/suscripciones/{id}", servidor.BorrarSuscripcion)

			r.Get("/accesos", servidor.ListarAccesos)
			r.Post("/accesos", servidor.CrearAcceso)
			r.Get("/accesos/{id}", servidor.ObtenerAcceso)
			r.Put("/accesos/{id}", servidor.ActualizarAcceso)
			r.With(soloAdmin).Delete("/accesos/{id}", servidor.BorrarAcceso)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
