// Package middleware contiene middleware HTTP compartido entre handlers.
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const bearerContextKey = "bearer_token"

// RequireAuth exige un header Authorization: Bearer <token>.
//
// Es un placeholder estructural (T-06 del backlog de Sprint 0): valida el
// formato del header, pero todavía no verifica firma ni expiración de un
// JWT real, porque EP-01 (US-002, login) aún no emite tokens. Se completa
// cuando se implemente esa historia.
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := extractBearerToken(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing_or_invalid_token"})
			return
		}
		c.Set(bearerContextKey, token)
		c.Next()
	}
}

// OptionalAuth no bloquea la request si falta el token (lo usa GET
// /ranking, que acepta tanto usuarios autenticados como anónimos), pero
// deja el token disponible en el contexto si vino.
func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, ok := extractBearerToken(c); ok {
			c.Set(bearerContextKey, token)
		}
		c.Next()
	}
}

func extractBearerToken(c *gin.Context) (string, bool) {
	const prefix = "Bearer "
	header := c.GetHeader("Authorization")
	if !strings.HasPrefix(header, prefix) {
		return "", false
	}
	token := strings.TrimPrefix(header, prefix)
	if token == "" {
		return "", false
	}
	return token, true
}
