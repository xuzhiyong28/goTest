package gormStudy

import (
	"time"
)

type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID
	Code string `json:"id" gorm:"index"`
	Timestamps
}
