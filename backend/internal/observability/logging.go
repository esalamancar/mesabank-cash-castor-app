// Package observability agrupa logging estructurado y métricas (T-11 del
// backlog de Sprint 0).
package observability

import (
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

// NewLogger crea un logger JSON a stdout (apto para recolectores de logs
// en Kubernetes).
func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

// RequestLogger loguea cada request con su request_id (tomado del header
// X-Request-Id si viene, o generado si no) y, cuando aplique, el código de
// partida (:code) para poder correlacionar por game_id.
func RequestLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-Id")
		if requestID == "" {
			requestID = uuid.NewString()
		}
		c.Set(RequestIDKey, requestID)
		c.Writer.Header().Set("X-Request-Id", requestID)

		start := time.Now()
		c.Next()

		fields := []any{
			"request_id", requestID,
			"method", c.Request.Method,
			"path", c.FullPath(),
			"status", c.Writer.Status(),
			"duration_ms", time.Since(start).Milliseconds(),
		}
		if code := c.Param("code"); code != "" {
			fields = append(fields, "game_code", code)
		}

		logger.Info("http_request", fields...)
	}
}
