package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	From      uuid.UUID `json:"from,omitempty"`
	To        uuid.UUID `json:"to,omitempty"`
	Content   string    `json:"content,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
