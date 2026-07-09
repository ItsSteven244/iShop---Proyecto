package models

import "time"

type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"unique;not null"`
	PasswordHash string    `json:"password" gorm:"not null"`
	Rol          string    `json:"rol" gorm:"not null;default:tecnico"`
	CreadoEn     time.Time `json:"creado_en"`
}
