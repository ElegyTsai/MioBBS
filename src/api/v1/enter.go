package v1

import (
	"Miyo_Assignment/api/v1/system"
	"Miyo_Assignment/service"
)

type ApiGroup struct {
	systemApi.SystemApi
}

var ApiGroupApp = new(ApiGroup)

var (
	UserServiceApp   = service.ServiceGroupApp.UserService
	CasbinServiceApp = service.ServiceGroupApp.CasbinService
)
