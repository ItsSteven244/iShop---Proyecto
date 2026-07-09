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

Hay dos roles: **admin** y **tecnico**. Cualquiera de los dos puede leer, crear y actualizar. **Solo admin puede borrar** (todos los endpoints `DELETE` de los 3 módulos están restringidos a ese rol).

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/auth/register | Registrar usuario (`email`, `password`, `rol` opcional — default `tecnico`) |
| POST | /api/v1/auth/login | Login, devuelve el JWT |

## Endpoints — Módulo Correctivo (Steven Soledispa)

### Ordenes Correctivas
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/correctivos/ordenes | Crear orden correctiva |
| GET | /api/v1/correctivos/ordenes | Listar ordenes correctivas |
| GET | /api/v1/correctivos/ordenes/{id} | Obtener orden por ID |
| PUT | /api/v1/correctivos/ordenes/{id} | Actualizar orden completa |
| PATCH | /api/v1/correctivos/ordenes/{id} | Actualizar estado y diagnostico |
| DELETE | /api/v1/correctivos/ordenes/{id} | Eliminar orden (solo admin) |

### Procesos de Reparacion
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/correctivos/procesos | Crear proceso de reparacion |
| GET | /api/v1/correctivos/procesos | Listar procesos |
| GET | /api/v1/correctivos/procesos/{id} | Obtener proceso por ID |
| PUT | /api/v1/correctivos/procesos/{id} | Actualizar proceso |
| DELETE | /api/v1/correctivos/procesos/{id} | Eliminar proceso (solo admin) |

### Evidencias de Daño
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/correctivos/evidencias | Crear evidencia de daño |
| GET | /api/v1/correctivos/evidencias | Listar evidencias |
| GET | /api/v1/correctivos/evidencias/{id} | Obtener evidencia por ID |
| PUT | /api/v1/correctivos/evidencias/{id} | Actualizar evidencia |
| DELETE | /api/v1/correctivos/evidencias/{id} | Eliminar evidencia (solo admin) |

## Endpoints — Módulo Preventivo (Ronnie Mera)

### Mantenimientos
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/mantenimientos | Crear mantenimiento preventivo |
| GET | /api/v1/mantenimientos | Listar mantenimientos |
| GET | /api/v1/mantenimientos/{id} | Obtener mantenimiento por ID |
| PUT | /api/v1/mantenimientos/{id} | Actualizar mantenimiento |
| DELETE | /api/v1/mantenimientos/{id} | Eliminar mantenimiento (solo admin) |

### Tareas Preventivas
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/tareas | Crear tarea preventiva |
| GET | /api/v1/tareas | Listar tareas |
| GET | /api/v1/tareas/{id} | Obtener tarea por ID |
| PUT | /api/v1/tareas/{id} | Actualizar tarea |
| DELETE | /api/v1/tareas/{id} | Eliminar tarea (solo admin) |

### Insumos Preventivos
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/insumos | Crear insumo |
| GET | /api/v1/insumos | Listar insumos |
| GET | /api/v1/insumos/{id} | Obtener insumo por ID |
| PUT | /api/v1/insumos/{id} | Actualizar insumo |
| DELETE | /api/v1/insumos/{id} | Eliminar insumo (solo admin) |

## Endpoints — Módulo Suscripciones y Servicios Digitales (Luis Litardo)

### Servicios Digitales
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/servicios | Crear servicio digital |
| GET | /api/v1/servicios | Listar servicios |
| GET | /api/v1/servicios/{id} | Obtener servicio por ID |
| PUT | /api/v1/servicios/{id} | Actualizar servicio |
| DELETE | /api/v1/servicios/{id} | Eliminar servicio (solo admin) |

### Suscripciones de Cliente
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/suscripciones | Crear suscripción |
| GET | /api/v1/suscripciones | Listar suscripciones |
| GET | /api/v1/suscripciones/{id} | Obtener suscripción por ID |
| PUT | /api/v1/suscripciones/{id} | Actualizar suscripción |
| DELETE | /api/v1/suscripciones/{id} | Eliminar suscripción (solo admin) |

### Accesos Digitales
| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /api/v1/accesos | Crear acceso digital |
| GET | /api/v1/accesos | Listar accesos |
| GET | /api/v1/accesos/{id} | Obtener acceso por ID |
| PUT | /api/v1/accesos/{id} | Actualizar acceso |
| DELETE | /api/v1/accesos/{id} | Eliminar acceso (solo admin) |

## Tests

Cada módulo tiene su suite de tests unitarios con mocks vía testify (service, handler con httptest, y repository con SQLite in-memory). Para correrlos:

```bash
go test ./... -v
```