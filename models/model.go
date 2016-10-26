package models

import "time"

//Base base model
type Base struct {
	ID        uint       `json:"id"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}
