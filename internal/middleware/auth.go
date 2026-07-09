package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ItsSteven244/iShop---Proyecto/internal/service"
)

type claveContext string

const (
	ClaveUsuarioID claveContext = "usuarioID"
	ClaveRol       claveContext = "rol"
)

func Auth(auth *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || partes[0] != "Bearer" {
				responderNoAutorizado(w)
				return
			}
			usuarioID, rol, err := auth.ValidarToken(partes[1])
			if err != nil {
				responderNoAutorizado(w)
				return
			}
			ctx := context.WithValue(r.Context(), ClaveUsuarioID, usuarioID)
			ctx = context.WithValue(ctx, ClaveRol, rol)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequireRol exige que el usuario autenticado tenga uno de los roles dados.
// Debe usarse DESPUÉS de Auth en la cadena de middlewares.
func RequireRol(rolesPermitidos ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rol, _ := r.Context().Value(ClaveRol).(string)
			for _, permitido := range rolesPermitidos {
				if rol == permitido {
					next.ServeHTTP(w, r)
					return
				}
			}
			responderProhibido(w)
		})
	}
}

func responderNoAutorizado(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"error": "token inexistente o invalido"}`))
}

func responderProhibido(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(`{"error": "no tienes permiso para esta accion"}`))
}
