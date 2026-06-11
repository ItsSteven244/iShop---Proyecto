package models

// ServicioDigital representa un servicio como Netflix, Spotify, etc.
// Guarda el nombre, categoría, precio, duración y cuántos perfiles incluye.
type ServicioDigital struct {
	ID               int     `json:"id"`
	Nombre           string  `json:"nombre"`
	Categoria        string  `json:"categoria"`
	Precio           float64 `json:"precio"`
	DuracionDias     int     `json:"duracion_dias"`
	CantidadPerfiles int     `json:"cantidad_perfiles"`
}

// SuscripcionCliente vincula a un cliente con un servicio digital.
// Guarda las fechas de inicio y fin, el estado y quién la gestionó.
type SuscripcionCliente struct {
	ID                int    `json:"id"`
	FechaInicio       string `json:"fecha_inicio"`
	FechaFin          string `json:"fecha_fin"`
	Estado            string `json:"estado"`
	ClienteID         int    `json:"cliente_id"`
	ServicioDigitalID int    `json:"servicio_digital_id"`
	TecnicoID         int    `json:"tecnico_id"`
}

// AccesoDigital guarda las credenciales de acceso de una suscripción activa.
// Está ligado a una suscripción mediante SuscripcionClienteID.
type AccesoDigital struct {
	ID                   int    `json:"id"`
	CorreoAcceso         string `json:"correo_acceso"`
	Perfil               string `json:"perfil"`
	Estado               string `json:"estado"`
	SuscripcionClienteID int    `json:"suscripcion_cliente_id"`
}
