package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"                  //Importa el router chi
	chimw "github.com/go-chi/chi/v5/middleware" //antes de una peticion pasa por filtros

	"github.com/ItsSteven244/iShop---Proyecto/internal/handlers"
	"github.com/ItsSteven244/iShop---Proyecto/internal/storage"
)

func main() {
	// 1. Crear el storage del módulo correctivo / tanto para ordenes, procesos y evidencias.
	correctivoStorage := storage.NuevaMemoriaCorrectivo()

	// 2. Configurar el router principal.
	r := chi.NewRouter()   //Se crea el router vacio, listo para registros
	r.Use(chimw.Logger)    //como un registro, imprime mensaje/logs
	r.Use(chimw.Recoverer) //si huo un problema no cae el server

	// 3. Montar el subrouter del módulo correctivo.
	r.Mount("/api/v1/correctivos", handlers.CorrectivosRouter(correctivoStorage))

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
