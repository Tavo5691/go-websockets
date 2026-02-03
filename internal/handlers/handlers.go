package handlers

import "gochat/internal/hub"

type Handler struct {
	hub    *hub.Hub
	jwtKey string
}

func New(h *hub.Hub, jwtKey string) *Handler {
	return &Handler{hub: h, jwtKey: jwtKey}
}
