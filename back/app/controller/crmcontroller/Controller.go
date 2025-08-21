package crmcontroller

import "tenjin/back/internal/crm"

type CrmController struct {
	userService *crm.UserService
}

func NewCrmController(userService *crm.UserService) *CrmController {
	return &CrmController{
		userService: userService,
	}
}
