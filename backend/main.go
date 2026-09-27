package main

import (
	"flag"
	"log"

	"github.com/esalamancar/mesabank-cash-castor-app/backend/internal/api"
	"github.com/esalamancar/mesabank-cash-castor-app/backend/internal/db"
	"github.com/esalamancar/mesabank-cash-castor-app/backend/internal/models"
	"github.com/esalamancar/mesabank-cash-castor-app/backend/internal/ws"
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

	// La conexión a datos para los handlers reales se agrega historia por
	// historia a partir de Sprint 1; por ahora todos son stubs (T-06).
	hub := ws.NewHub()
	router := api.NewRouter(hub)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("el servidor terminó con error: %v", err)
	}
}
