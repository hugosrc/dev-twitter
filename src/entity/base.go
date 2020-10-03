package entity

import (
	"time"
)

// Base is the model responsible for setting the default values
type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
