package models

// OrdenCorrectiva es la orden que se crea cuando llega un dispositivo a reparar.
// Guarda el problema que reportó el cliente, el diagnóstico del técnico y el estado actual.
type OrdenCorrectiva struct {
	ID                int     `json:"id"`
	Codigo            string  `json:"codigo"`
	ProblemaReportado string  `json:"problema_reportado"`
	Diagnostico       string  `json:"diagnostico"`
	Estado            string  `json:"estado"`
	Prioridad         string  `json:"prioridad"`
	Costo             float64 `json:"costo"`
	FechaIngreso      string  `json:"fecha_ingreso"`
	FechaEntrega      string  `json:"fecha_entrega"`
	DispositivoID     int     `json:"dispositivo_id"`
	TecnicoID         int     `json:"tecnico_id"`
}

// ProcesoReparacion representa cada paso que se hace durante la reparación.
// Está ligado a una orden correctiva mediante OrdenCorrectivaID.
type ProcesoReparacion struct {
	ID                int    `json:"id"`
	Etapa             string `json:"etapa"`
	Observacion       string `json:"observacion"`
	Fecha             string `json:"fecha"`
	OrdenCorrectivaID int    `json:"orden_correctiva_id"`
}

// EvidenciaDanio guarda la descripción o foto del daño que tiene el dispositivo.
// También va ligada a una orden correctiva.
type EvidenciaDanio struct {
	ID                int    `json:"id"`
	Descripcion       string `json:"descripcion"`
	ImagenURL         string `json:"imagen_url"`
	Fecha             string `json:"fecha"`
	OrdenCorrectivaID int    `json:"orden_correctiva_id"`
}
