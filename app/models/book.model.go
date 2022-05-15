package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID      uuid.UUID
	Title   string
	Author  string
	Summary string
}
