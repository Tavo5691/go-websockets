package handlers

import "gochat/internal/hub"

type Handler struct {
	hub *hub.Hub
}

func New(h *hub.Hub) *Handler {
	return &Handler{hub: h}
}
