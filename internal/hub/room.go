package hub

import "github.com/google/uuid"

type Room struct {
	roomId  uuid.UUID
	members map[*Client]bool
}

func NewRoom(roomId uuid.UUID) *Room {
	return &Room{
		roomId:  roomId,
		members: make(map[*Client]bool),
	}
}
