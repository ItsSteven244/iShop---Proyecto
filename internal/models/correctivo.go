package models

// OrdenCorrectiva representa una orden de reparación correctiva de un dispositivo.
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

// ProcesoReparacion representa una etapa del proceso de reparación de una orden correctiva.
type ProcesoReparacion struct {
	ID                int    `json:"id"`
	Etapa             string `json:"etapa"`
	Observacion       string `json:"observacion"`
	Fecha             string `json:"fecha"`
	OrdenCorrectivaID int    `json:"orden_correctiva_id"`
}

// EvidenciaDanio representa una evidencia fotográfica o descriptiva del daño reportado.
type EvidenciaDanio struct {
	ID                int    `json:"id"`
	Descripcion       string `json:"descripcion"`
	ImagenURL         string `json:"imagen_url"`
	Fecha             string `json:"fecha"`
	OrdenCorrectivaID int    `json:"orden_correctiva_id"`
}
