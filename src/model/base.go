package model

import "time"

// Base contains base model info
type Base struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
