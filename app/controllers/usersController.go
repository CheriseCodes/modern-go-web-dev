package controllers

type UsersController struct {
	usersService *services.UsersService
}

func NewUsersController(
	usersService *services.UsersService
) *UsersController {
	return &UsersController{
		usersService: usersService,
	}
}

func (uc UsersController) Login(ctx *gin.Context) {
	// TODO
}

func (uc UsersController) Logout(ctx *gin.Context) {
	// TODO
}


