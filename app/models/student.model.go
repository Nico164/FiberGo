package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID      uuid.UUID
	Name    string
	Address string
	Phone   string
	Email   string
	Gender  string
	Born    string
}
