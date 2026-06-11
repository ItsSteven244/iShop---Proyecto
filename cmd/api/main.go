package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"github.com/ItsSteven244/iShop---Proyecto/internal/handlers"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

func main() {
	// 1. Crear el storage del módulo correctivo.
	correctivoStorage := storage.NuevaMemoriaCorrectivo()

	// 2. Crear el storage del módulo de suscripciones.
	suscripcionesStorage := storage.NuevaMemoriaSuscripciones()

	// 3. Configurar el router principal.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	// 4. Montar el subrouter del módulo correctivo.
	r.Mount("/api/v1/correctivos", handlers.CorrectivosRouter(correctivoStorage))

	// 5. Montar el subrouter del módulo de suscripciones.
	r.Mount("/api/v1/suscripciones", handlers.SuscripcionesRouter(suscripcionesStorage))

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
