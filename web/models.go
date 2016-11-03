package web

import "time"

//Model basic db model
type Model struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
