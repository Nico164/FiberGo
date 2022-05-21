package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID      uuid.UUID
	Title   string `json:"title"`
	Author  string `json:"author"`
	Summary string `json:"summary"`
}
