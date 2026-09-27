// Package models define el esquema de datos (PRD §5) como modelos GORM.
// El esquema se aplica con AutoMigrate (ADR-008), no con SQL manual.
package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username  string    `gorm:"size:32;unique;not null"`
	PinHash   string    `gorm:"size:255;not null"`
	IsGuest   bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

type Game struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Code      string    `gorm:"size:16;unique;not null"`
	QRURL     string    `gorm:"column:qr_url"`
	Link      string
	Status    string    `gorm:"size:16;not null;default:lobby"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt time.Time `gorm:"not null;default:now()"`

	Config *GameConfig `gorm:"foreignKey:GameID"`
}

// GameConfig es 1:1 con Game. El campo Inflation existe pero no tiene
// efecto funcional en el MVP (decisión del PO, ver US-036/US-201).
type GameConfig struct {
	GameID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Currency          string    `gorm:"size:8;not null"`
	ExchangeRate      float64   `gorm:"not null;default:1"`
	Inflation         *float64
	Fee               float64 `gorm:"not null;default:0"`
	InterestType      *string `gorm:"size:16"`
	InterestRate      *float64
	InterestInterval  *int
	MaxDebtMultiplier float64
	TotalDigital      float64 `gorm:"not null"`
	TotalPaper        float64 `gorm:"not null"`
	InitialPerPlayer  float64 `gorm:"not null"`
}

type Player struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	GameID         uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_players_game_user"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_players_game_user"`
	Role           string    `gorm:"size:16;not null"`
	DigitalBalance float64   `gorm:"not null;default:0"`
	PaperBalance   float64   `gorm:"not null;default:0"`
	Debt           float64   `gorm:"not null;default:0"`
	IsBankrupt     bool      `gorm:"not null;default:false"`
}

type Transaction struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	GameID         uuid.UUID  `gorm:"type:uuid;not null;index"`
	FromPlayerID   *uuid.UUID `gorm:"type:uuid"`
	ToPlayerID     *uuid.UUID `gorm:"type:uuid"`
	Amount         float64    `gorm:"not null"`
	Type           string     `gorm:"size:32;not null"`
	IdempotencyKey *uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	CreatedAt      time.Time  `gorm:"not null;default:now()"`
}

type Loan struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	GameID       uuid.UUID `gorm:"type:uuid;not null;index"`
	PlayerID     uuid.UUID `gorm:"type:uuid;not null;index"`
	Amount       float64   `gorm:"not null"`
	InterestRate float64   `gorm:"not null"`
	DueRound     *int
	Status       string `gorm:"size:16;not null;default:active"`
}

type AuditLog struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	GameID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	PlayerID  *uuid.UUID `gorm:"type:uuid"`
	Action    string     `gorm:"size:64;not null"`
	Details   string
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

// All lista todos los modelos, para AutoMigrate y para cualquier otro
// código que necesite iterarlos.
func All() []interface{} {
	return []interface{}{
		&User{}, &Game{}, &GameConfig{}, &Player{}, &Transaction{}, &Loan{}, &AuditLog{},
	}
}
