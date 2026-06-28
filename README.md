# API Backend - iShop

Proyecto backend desarrollado en Go para la gestión de mantenimiento y servicios digitales de dispositivos móviles en iShop.

## Módulos
- Mantenimiento Preventivo
- Reparación Correctiva
- Suscripciones y Servicios Digitales

## Tecnologías
- Go
- PostgreSQL
- JWT
- Docker

## Integrantes
- Luis Litardo
- Steven Soledispa
- Ronnie Mera

## Endpoints — Módulo Correctivo (Steven Soledispa)

### Ordenes Correctivas
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/correctivos/ordenes | Crear orden correctiva |
| GET | /api/v1/correctivos/ordenes | Listar ordenes correctivas |
| GET | /api/v1/correctivos/ordenes/{id} | Obtener orden por ID |
| PUT | /api/v1/correctivos/ordenes/{id} | Actualizar orden completa |
| PATCH | /api/v1/correctivos/ordenes/{id} | Actualizar estado y diagnostico |
| DELETE | /api/v1/correctivos/ordenes/{id} | Eliminar orden |

### Procesos de Reparacion
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/correctivos/procesos | Crear proceso de reparacion |
| GET | /api/v1/correctivos/procesos | Listar procesos |
| GET | /api/v1/correctivos/procesos/{id} | Obtener proceso por ID |
| PUT | /api/v1/correctivos/procesos/{id} | Actualizar proceso |
| DELETE | /api/v1/correctivos/procesos/{id} | Eliminar proceso |

### Evidencias de Daño
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/correctivos/evidencias | Crear evidencia de daño |
| GET | /api/v1/correctivos/evidencias | Listar evidencias |
| GET | /api/v1/correctivos/evidencias/{id} | Obtener evidencia por ID |
| PUT | /api/v1/correctivos/evidencias/{id} | Actualizar evidencia |
| DELETE | /api/v1/correctivos/evidencias/{id} | Eliminar evidencia |