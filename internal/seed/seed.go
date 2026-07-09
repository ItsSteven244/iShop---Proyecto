package seed

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
)

// Ejecutar siembra datos de ejemplo si la base está vacía.
// Se llama una sola vez al arrancar el servidor (main.go), después de AutoMigrate.
// No hace nada si ya existen usuarios (evita duplicar datos en cada reinicio).
func Ejecutar(db *gorm.DB) error {
	var totalUsuarios int64
	if err := db.Model(&models.Usuario{}).Count(&totalUsuarios).Error; err != nil {
		return err
	}
	if totalUsuarios > 0 {
		log.Println("Seed: ya hay datos, se omite la siembra.")
		return nil
	}

	log.Println("Seed: base vacía, sembrando datos de ejemplo...")

	hashAdmin, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashTecnico, err := bcrypt.GenerateFromPassword([]byte("tecnico123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := models.Usuario{Email: "admin@ishop.com", PasswordHash: string(hashAdmin), Rol: "admin"}
	tecnico := models.Usuario{Email: "tecnico@ishop.com", PasswordHash: string(hashTecnico), Rol: "tecnico"}
	if err := db.Create(&admin).Error; err != nil {
		return err
	}
	if err := db.Create(&tecnico).Error; err != nil {
		return err
	}

	// Módulo Correctivo
	orden := models.OrdenCorrectiva{
		Codigo:            "ORD-001",
		ProblemaReportado: "Pantalla rota",
		Diagnostico:       "Reemplazo de pantalla necesario",
		Estado:            "en_proceso",
		Prioridad:         "alta",
		Costo:             45.50,
		FechaIngreso:      "2026-07-01",
		FechaEntrega:      "2026-07-05",
		DispositivoID:     1,
		TecnicoID:         tecnico.ID,
	}
	if err := db.Create(&orden).Error; err != nil {
		return err
	}
	proceso := models.ProcesoReparacion{
		Etapa:             "diagnostico",
		Observacion:       "Se confirma daño en pantalla",
		Fecha:             "2026-07-01",
		OrdenCorrectivaID: orden.ID,
	}
	if err := db.Create(&proceso).Error; err != nil {
		return err
	}
	evidencia := models.EvidenciaDanio{
		Descripcion:       "Foto de la pantalla rota",
		ImagenURL:         "https://ejemplo.com/foto1.jpg",
		Fecha:             "2026-07-01",
		OrdenCorrectivaID: orden.ID,
	}
	if err := db.Create(&evidencia).Error; err != nil {
		return err
	}

	// Módulo Preventivo
	mantenimiento := models.MantenimientoPreventivo{
		Equipo:            "Laptop Dell XPS",
		FechaProgramada:   "2026-07-10",
		TipoMantenimiento: "Limpieza interna",
		Estado:            "pendiente",
		TecnicoID:         tecnico.ID,
	}
	if err := db.Create(&mantenimiento).Error; err != nil {
		return err
	}
	tarea := models.TareaPreventiva{
		Descripcion:               "Limpiar ventiladores",
		Estado:                    "pendiente",
		Duracion:                  30,
		MantenimientoPreventivoID: mantenimiento.ID,
	}
	if err := db.Create(&tarea).Error; err != nil {
		return err
	}
	insumo := models.InsumoPreventivo{
		Nombre:                    "Pasta térmica",
		Cantidad:                  1,
		Costo:                     3.50,
		MantenimientoPreventivoID: mantenimiento.ID,
	}
	if err := db.Create(&insumo).Error; err != nil {
		return err
	}

	// Módulo Suscripciones
	servicio := models.ServicioDigital{
		Nombre:           "Netflix Premium",
		Categoria:        "Streaming",
		Precio:           9.99,
		DuracionDias:     30,
		CantidadPerfiles: 4,
	}
	if err := db.Create(&servicio).Error; err != nil {
		return err
	}
	suscripcion := models.SuscripcionCliente{
		FechaInicio:       "2026-07-01",
		FechaFin:          "2026-07-31",
		Estado:            "activa",
		ClienteID:         1,
		ServicioDigitalID: servicio.ID,
		TecnicoID:         tecnico.ID,
	}
	if err := db.Create(&suscripcion).Error; err != nil {
		return err
	}
	acceso := models.AccesoDigital{
		CorreoAcceso:         "cliente1@correo.com",
		Perfil:               "Perfil 1",
		Estado:               "activo",
		SuscripcionClienteID: suscripcion.ID,
	}
	if err := db.Create(&acceso).Error; err != nil {
		return err
	}

	log.Println("Seed: datos de ejemplo creados correctamente.")
	return nil
}
