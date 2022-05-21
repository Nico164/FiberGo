package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID     uuid.UUID
	Name   string `json:"name"`
	Major  string `json:"major"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}
