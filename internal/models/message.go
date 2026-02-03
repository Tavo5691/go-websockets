package models

import (
	"time"

	"github.com/google/uuid"
)

type MessageType string

const (
	Room  MessageType = "room"
	Join  MessageType = "join"
	Leave MessageType = "leave"
	DM    MessageType = "dm"
)

type Message struct {
	Type      MessageType `json:"type,omitempty"`
	Room      uuid.UUID   `json:"room,omitempty"`
	From      uuid.UUID   `json:"from,omitempty"`
	To        uuid.UUID   `json:"to,omitempty"`
	Content   string      `json:"content,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
}
