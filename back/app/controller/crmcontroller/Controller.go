package crmcontroller

import (
	"tenjin/back/internal/crm"
	"tenjin/back/internal/emailverification"
)

type CrmController struct {
	userService *crm.UserService
	emailVerificationService *emailverification.EmailVerificationService
}

func NewCrmController(userService *crm.UserService, emailVerificationUser *emailverification.EmailVerificationService) *CrmController {
	return &CrmController{
		userService: userService,
		emailVerificationService: emailVerificationUser,
	}
}
