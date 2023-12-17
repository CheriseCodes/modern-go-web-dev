package services

type UsersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersService(
	usersRepository *repositories.UsersRepository
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository
	}
}

func (us UsersService) Login(username string, password string) (string, *models.ResponseError) {
	// TODO
}

func (us UsersService) Logout(accessToken string) *models.ResponseError  {
	// TODO
}

func (us UsersService) AuthorizeUser(accessToken string, expectedRoles []string) (bool, *models.ResponseError) {
	// TODO
}

func generateAccessToken(username string) (string, *models.ResponseError) {
	// TODO
}
