package handlers

import (
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

type Server struct {
	// Módulo Correctivo - Steven
	Ordenes    *service.OrdenCorrectivaService
	Procesos   *service.ProcesoReparacionService
	Evidencias *service.EvidenciaDanioService

	// Auth
	Auth *service.AuthService

	// Módulo Preventivo - (compañero agrega aquí)

	// Módulo Suscripciones - (compañero agrega aquí)
}

func NewServer(
	ordenes *service.OrdenCorrectivaService,
	procesos *service.ProcesoReparacionService,
	evidencias *service.EvidenciaDanioService,
	auth *service.AuthService,
) *Server {
	return &Server{
		Ordenes:    ordenes,
		Procesos:   procesos,
		Evidencias: evidencias,
		Auth:       auth,
	}
}
