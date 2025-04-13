package api

import "avitopvz/internal/service"

type UserHandler struct {
	service service.Interface
}

func NewUserHandler(service service.Interface) *UserHandler {
	return &UserHandler{service: service}
}
