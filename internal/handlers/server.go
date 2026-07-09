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

	// Módulo Preventivo
	Mantenimiento *service.MantenimientoPreventivoService
	Tareas        *service.TareaPreventivaService
	Insumos       *service.InsumoPreventivoService

	// Módulo Suscripciones - Luisao
	Servicios     *service.ServicioDigitalService
	Suscripciones *service.SuscripcionClienteService
	Accesos       *service.AccesoDigitalService
}

func NewServer(
	ordenes *service.OrdenCorrectivaService,
	procesos *service.ProcesoReparacionService,
	evidencias *service.EvidenciaDanioService,
	auth *service.AuthService,
	servicios *service.ServicioDigitalService,
	suscripciones *service.SuscripcionClienteService,
	accesos *service.AccesoDigitalService,
	mantenimiento *service.MantenimientoPreventivoService,
	tareas *service.TareaPreventivaService,
	insumos *service.InsumoPreventivoService,
) *Server {
	return &Server{
		Ordenes:       ordenes,
		Procesos:      procesos,
		Evidencias:    evidencias,
		Auth:          auth,
		Servicios:     servicios,
		Suscripciones: suscripciones,
		Accesos:       accesos,
		Mantenimiento: mantenimiento,
		Tareas:        tareas,
		Insumos:       insumos,
	}
}
