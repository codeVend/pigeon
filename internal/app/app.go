package app

import (
	"github.com/pigeon/internal/auth"
	"pigeon/internal/http"
)

func Run() error {
	authService := &auth.Service{}
	handler := &http.Handler{AuthService: authService}

	testUser := auth.User{
		ID:       1,
		Username: "asdasdas",
		Email:    "danil@bad.org",
		Password: "1234",
	}

	return handler.HandleRegister(testUser)
}
