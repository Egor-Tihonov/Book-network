// Package handler ...
package handler

import (
	"github.com/Egor-Tihonov/Book-network/internal/model"
	"github.com/Egor-Tihonov/Book-network/internal/service"
)

// Handler ...
type Handler struct {
	se *service.Service
	model.MyCookie
}

// New create new handler
func New(srv *service.Service, c model.MyCookie) *Handler {
	return &Handler{se: srv, MyCookie: c}
}
