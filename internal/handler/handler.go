// Package handler ...
package handler

import (
	"github.com/Egor-Tihonov/Book-network/internal/service"
)

// Handler ...
type Handler struct {
	se *service.Service
}

// New create new handler
func New(srv *service.Service) *Handler {
	return &Handler{se: srv}
}
