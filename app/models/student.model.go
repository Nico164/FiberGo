package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID      uuid.UUID
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
	Born    string `json:"born"`
}
