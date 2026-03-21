package http

import (
		"pigeon/internal/auth"
		"net/http"
	)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	

	type Handler struct {
		AuthService *auth.Service
	}

	func (h *Handler) HandleRegister(u *auth.User); error {
		return h.AuthService.Register(u)
	}

}

