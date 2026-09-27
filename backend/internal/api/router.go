// Package api arma el router HTTP: un handler stub por cada operación de
// docs/08-openapi.yaml (T-06 del backlog de Sprint 0). La lógica real de
// cada endpoint se implementa historia por historia (docs/07-historias-usuario.md)
// a partir de Sprint 1.
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/esalamancar/mesabank-cash-castor-app/backend/internal/middleware"
	"github.com/esalamancar/mesabank-cash-castor-app/backend/internal/ws"
)

// stub responde 501 identificando el endpoint, para poder verificar el
// contrato de rutas contra el OpenAPI antes de que exista lógica real.
func stub(endpoint string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"error":    "not_implemented",
			"endpoint": endpoint,
		})
	}
}

// NewRouter arma todas las rutas del contrato. hub es el hub de WebSocket
// (paquete ws) compartido por toda la app.
func NewRouter(hub *ws.Hub) *gin.Engine {
	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", stub("POST /auth/register"))
		authGroup.POST("/login", stub("POST /auth/login"))
		authGroup.POST("/guest", stub("POST /auth/guest"))
		authGroup.POST("/refresh", stub("POST /auth/refresh"))
		authGroup.POST("/logout", middleware.RequireAuth(), stub("POST /auth/logout"))
	}

	router.GET("/users/me", middleware.RequireAuth(), stub("GET /users/me"))

	gamesGroup := router.Group("/games", middleware.RequireAuth())
	{
		gamesGroup.POST("", stub("POST /games"))
		gamesGroup.GET("/:code", stub("GET /games/{code}"))
		gamesGroup.POST("/:code/join", stub("POST /games/{code}/join"))
		gamesGroup.GET("/:code/players", stub("GET /games/{code}/players"))
		gamesGroup.POST("/:code/players/:playerId/expel", stub("POST /games/{code}/players/{playerId}/expel"))
		gamesGroup.POST("/:code/reset", stub("POST /games/{code}/reset"))
		gamesGroup.POST("/:code/reset/confirm", stub("POST /games/{code}/reset/confirm"))

		gamesGroup.GET("/:code/balance", stub("GET /games/{code}/balance"))
		gamesGroup.POST("/:code/transfer", stub("POST /games/{code}/transfer"))
		gamesGroup.POST("/:code/deposit", stub("POST /games/{code}/deposit"))
		gamesGroup.POST("/:code/loans", stub("POST /games/{code}/loans"))
		gamesGroup.POST("/:code/loans/:loanId/repay", stub("POST /games/{code}/loans/{loanId}/repay"))
		gamesGroup.POST("/:code/players/:playerId/liquidate", stub("POST /games/{code}/players/{playerId}/liquidate"))
		gamesGroup.GET("/:code/cash-audit", stub("GET /games/{code}/cash-audit"))
		gamesGroup.POST("/:code/cash-audit", stub("POST /games/{code}/cash-audit"))
		gamesGroup.GET("/:code/transactions", stub("GET /games/{code}/transactions"))
		gamesGroup.GET("/:code/audit", stub("GET /games/{code}/audit"))

		gamesGroup.POST("/:code/turn/next", stub("POST /games/{code}/turn/next"))
		gamesGroup.GET("/:code/ws", ws.ServeHTTP(hub))

		gamesGroup.POST("/:code/finish", stub("POST /games/{code}/finish"))
	}

	router.GET("/ranking", middleware.OptionalAuth(), stub("GET /ranking"))

	adminGroup := router.Group("/admin", middleware.RequireAuth())
	{
		adminGroup.GET("/currencies", stub("GET /admin/currencies"))
		adminGroup.POST("/currencies", stub("POST /admin/currencies"))
	}

	return router
}
