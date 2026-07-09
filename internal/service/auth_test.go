package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ItsSteven244/iShop---Proyecto/internal/models"
	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

// =========================================================
// FAKE REPO (en memoria) PARA LOS TESTS DE AUTH
// =========================================================

type fakeUserRepo struct {
	usuarios []models.Usuario
	nextID   int
}

func newFakeUserRepo() *fakeUserRepo {
	return &fakeUserRepo{nextID: 1}
}

func (f *fakeUserRepo) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.usuarios = append(f.usuarios, u)
	return u, nil
}

func (f *fakeUserRepo) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	for _, u := range f.usuarios {
		if u.Email == email {
			return u, true
		}
	}
	return models.Usuario{}, false
}

// =========================================================
// REGISTRAR
// =========================================================

func TestAuthService_Registrar_CredencialesVacias(t *testing.T) {
	// Preparar
	repo := newFakeUserRepo()
	auth := service.NewAuthService(repo)

	// Ejecutar — email vacío
	_, err := auth.Registrar("", "123456", "tecnico")

	// Verificar
	require.ErrorIs(t, err, service.ErrCredencialesInvalidas)
}

func TestAuthService_Registrar_RolInvalido(t *testing.T) {
	// Preparar
	repo := newFakeUserRepo()
	auth := service.NewAuthService(repo)

	// Ejecutar — rol que no es "admin" ni "tecnico"
	_, err := auth.Registrar("test@test.com", "123456", "superadmin")

	// Verificar
	require.ErrorIs(t, err, service.ErrRolInvalido)
}

func TestAuthService_Registrar_EmailEnUso(t *testing.T) {
	// Preparar — ya existe un usuario con ese email
	repo := newFakeUserRepo()
	auth := service.NewAuthService(repo)
	_, err := auth.Registrar("test@test.com", "123456", "tecnico")
	require.NoError(t, err)

	// Ejecutar — intentamos registrar el mismo email otra vez
	_, err = auth.Registrar("test@test.com", "otraPass", "tecnico")

	// Verificar
	require.ErrorIs(t, err, service.ErrEmailEnUso)
}

func TestAuthService_Registrar_Exito(t *testing.T) {
	// Preparar
	repo := newFakeUserRepo()
	auth := service.NewAuthService(repo)

	// Ejecutar — sin rol, debe asumir "tecnico" por defecto
	usuario, err := auth.Registrar("nuevo@test.com", "123456", "")

	// Verificar
	require.NoError(t, err)
	require.Equal(t, "nuevo@test.com", usuario.Email)
	require.Equal(t, "tecnico", usuario.Rol)
	require.NotEqual(t, "123456", usuario.PasswordHash) // debe estar hasheado, no en texto plano
}

// =========================================================
// LOGIN
// =========================================================

func TestAuthService_Login_UsuarioNoExiste(t *testing.T) {
	// Preparar
	repo := newFakeUserRepo()
	auth := service.NewAuthService(repo)

	// Ejecutar
	_, err := auth.Login("noexiste@test.com", "123456")

	// Verificar
	require.ErrorIs(t, err, service.ErrCredencialesInvalidas)
}

func TestAuthService_Login_PasswordIncorrecta(t *testing.T) {
	// Preparar
	repo := newFakeUserRepo()
	auth := service.NewAuthService(repo)
	_, err := auth.Registrar("test@test.com", "correcta123", "tecnico")
	require.NoError(t, err)

	// Ejecutar — password equivocada
	_, err = auth.Login("test@test.com", "incorrecta456")

	// Verificar
	require.ErrorIs(t, err, service.ErrCredencialesInvalidas)
}

func TestAuthService_Login_Exito(t *testing.T) {
	// Preparar
	repo := newFakeUserRepo()
	auth := service.NewAuthService(repo)
	_, err := auth.Registrar("test@test.com", "correcta123", "admin")
	require.NoError(t, err)

	// Ejecutar
	token, err := auth.Login("test@test.com", "correcta123")

	// Verificar — debe devolver un token JWT no vacío
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

// =========================================================
// VALIDAR TOKEN
// =========================================================

func TestAuthService_ValidarToken_Valido(t *testing.T) {
	// Preparar — registramos y logueamos para obtener un token real
	repo := newFakeUserRepo()
	auth := service.NewAuthService(repo)
	usuario, err := auth.Registrar("test@test.com", "correcta123", "admin")
	require.NoError(t, err)
	token, err := auth.Login("test@test.com", "correcta123")
	require.NoError(t, err)

	// Ejecutar
	usuarioID, rol, err := auth.ValidarToken(token)

	// Verificar
	require.NoError(t, err)
	require.Equal(t, usuario.ID, usuarioID)
	require.Equal(t, "admin", rol)
}

func TestAuthService_ValidarToken_Invalido(t *testing.T) {
	// Preparar
	repo := newFakeUserRepo()
	auth := service.NewAuthService(repo)

	// Ejecutar — token basura, no generado por nuestro servicio
	_, _, err := auth.ValidarToken("esto-no-es-un-token-valido")

	// Verificar
	require.ErrorIs(t, err, service.ErrCredencialesInvalidas)
}
