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

	// 2. Configurar el router principal.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	// 3. Montar el subrouter del módulo correctivo.
	r.Mount("/api/v1/correctivos", handlers.CorrectivosRouter(correctivoStorage))

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
