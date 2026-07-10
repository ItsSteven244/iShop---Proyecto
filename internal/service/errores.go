package service

import "errors"

var (
	ErrCodigoVacio           = errors.New("el campo codigo es obligatorio")
	ErrProblemaVacio         = errors.New("el campo problema_reportado es obligatorio")
	ErrEtapaVacia            = errors.New("el campo etapa es obligatorio")
	ErrDescripcionVacia      = errors.New("el campo descripcion es obligatorio")
	ErrOrdenIDRequerido      = errors.New("el campo orden_correctiva_id es obligatorio")
	ErrNoEncontrado          = errors.New("recurso no encontrado")
	ErrTokenInvalido         = errors.New("token inexistente o invalido")
	ErrEmailEnUso            = errors.New("email ya en uso")
	ErrCredencialesInvalidas = errors.New("email o contraseña incorrectos")
	ErrRolInvalido           = errors.New("el rol debe ser 'admin' o 'tecnico'")
	ErrNombreServicioVacio   = errors.New("el campo nombre es obligatorio")
	ErrPrecioInvalido        = errors.New("el campo precio debe ser mayor a 0")
	ErrEquipoVacio           = errors.New("el campo equipo es obligatorio")
	ErrFechaProgramadaVacia  = errors.New("el campo fecha_programada es obligatorio")

	// Suscripciones — SuscripcionCliente
	ErrClienteIDRequerido         = errors.New("el campo cliente_id es obligatorio")
	ErrServicioDigitalIDRequerido = errors.New("el campo servicio_digital_id es obligatorio")
	ErrFechaInicioVacia           = errors.New("el campo fecha_inicio es obligatorio")
	ErrFechaFinVacia              = errors.New("el campo fecha_fin es obligatorio")
	ErrEstadoVacio                = errors.New("el campo estado es obligatorio")
	ErrTecnicoIDRequerido         = errors.New("el campo tecnico_id es obligatorio")

	// Suscripciones — AccesoDigital.
	ErrCorreoAccesoVacio      = errors.New("el campo correo_acceso es obligatorio")
	ErrPerfilVacio            = errors.New("el campo perfil es obligatorio")
	ErrSuscripcionIDRequerido = errors.New("el campo suscripcion_cliente_id es obligatorio")

	// Preventivo — TareaPreventiva
	ErrDescripcionTareaVacia    = errors.New("el campo descripcion es obligatorio")
	ErrMantenimientoIDRequerido = errors.New("el campo mantenimiento_preventivo_id es obligatorio")

	// Preventivo — InsumoPreventivo
	ErrNombreInsumoVacio = errors.New("el campo nombre es obligatorio")
)
