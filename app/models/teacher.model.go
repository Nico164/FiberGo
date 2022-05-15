package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID     uuid.UUID
	Name   string
	Major  string
	Phone  string
	Email  string
	Gender string
}
