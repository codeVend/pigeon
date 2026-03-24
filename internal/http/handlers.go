package http

import (
		"github.com/codeVend/pigeon/internal/auth"
		"net/http"
	)
type NewHandler struct {
	AuthService *auth.Service
}

func (h *NewHandler) HandleRegister(u *auth.User) error {
	return h.AuthService.Register(u)
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	

	
}

