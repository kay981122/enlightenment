package system

import "em-app/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
