package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/esalamancar/mesabank-cash-castor-app/backend/internal/db"
	"github.com/esalamancar/mesabank-cash-castor-app/backend/internal/models"
)

func main() {
	migrationsOnly := flag.Bool("migrations-only", false, "corre AutoMigrate y termina, sin levantar el servidor (ADR-008)")
	flag.Parse()

	if *migrationsOnly {
		conn, err := db.Connect()
		if err != nil {
			log.Fatalf("no se pudo conectar a la base de datos: %v", err)
		}
		if err := conn.AutoMigrate(models.All()...); err != nil {
			log.Fatalf("fallo en AutoMigrate: %v", err)
		}
		log.Println("migraciones aplicadas correctamente")
		return
	}

	// La conexión a datos para el servidor HTTP se agrega en T-06, cuando
	// los handlers empiecen a necesitarla de verdad.
	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.Run(":8080")
}
