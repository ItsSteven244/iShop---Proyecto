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
	ErrNombreServicioVacio   = errors.New("el campo nombre es obligatorio")
	ErrPrecioInvalido        = errors.New("el campo precio debe ser mayor a 0")
)
