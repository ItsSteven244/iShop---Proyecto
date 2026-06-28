package models

// MantenimientoPreventivo representa una orden de mantenimiento preventivo programado.
// Guarda el equipo, la fecha programada, el tipo de mantenimiento y su estado.
type MantenimientoPreventivo struct {
	ID                int    `json:"id"`
	Equipo            string `json:"equipo"`
	FechaProgramada   string `json:"fecha_programada"`
	TipoMantenimiento string `json:"tipo_mantenimiento"`
	Estado            string `json:"estado"`
	TecnicoID         int    `json:"tecnico_id"`
}

// TareaPreventiva representa una tarea específica dentro de un mantenimiento preventivo.
// Está ligada a un mantenimiento mediante MantenimientoPreventivoID..
type TareaPreventiva struct {
	ID                        int    `json:"id"`
	Descripcion               string `json:"descripcion"`
	Estado                    string `json:"estado"`
	Duracion                  int    `json:"duracion_minutos"`
	MantenimientoPreventivoID int    `json:"mantenimiento_preventivo_id"`
}

// InsumoPreventivo representa un insumo o repuesto usado en un mantenimiento preventivo.
// Guarda el nombre, cantidad y a qué mantenimiento pertenece..
type InsumoPreventivo struct {
	ID                        int     `json:"id"`
	Nombre                    string  `json:"nombre"`
	Cantidad                  int     `json:"cantidad"`
	Costo                     float64 `json:"costo"`
	MantenimientoPreventivoID int     `json:"mantenimiento_preventivo_id"`
}
