package service

import "Miyo_Assignment/service/system"

type ServiceGroup struct {
	UserService   system.UserService
	CasbinService system.CasbinService
}

var ServiceGroupApp = new(ServiceGroup)
