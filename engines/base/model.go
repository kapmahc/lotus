package base

import "time"

//Model base model
type Model struct {
	ID        uint       `json:"id"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}
