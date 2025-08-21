package authcontroller

import (
	"tenjin/back/internal/auth"
	"tenjin/back/internal/crm"
)

type AuthController struct {
	userService *crm.UserService
	authService *auth.AuthService
}

func NewAuthController(userService *crm.UserService, authService *auth.AuthService) *AuthController {
	return &AuthController{
		userService: userService,
		authService: authService,
	}
}
