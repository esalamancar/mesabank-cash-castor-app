// Package db abre la conexión a PostgreSQL vía GORM.
package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Connect abre la conexión usando las variables de entorno documentadas
// en .env.example (POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER,
// POSTGRES_PASSWORD, POSTGRES_DB).
func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getenv("POSTGRES_HOST", "localhost"),
		getenv("POSTGRES_PORT", "5432"),
		getenv("POSTGRES_USER", "cashcastor"),
		getenv("POSTGRES_PASSWORD", "cashcastor"),
		getenv("POSTGRES_DB", "cashcastor"),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// Ping intenta conectar y hacer ping a la base de datos. Se usa en
// /readyz (T-11): a diferencia de Connect(), no se guarda la conexión —
// cada chequeo de readiness abre y cierra la suya para reflejar el
// estado real en el momento de la consulta.
func Ping() error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	sqlDB, err := conn.DB()
	if err != nil {
		return err
	}
	defer func() { _ = sqlDB.Close() }()

	return sqlDB.Ping()
}
