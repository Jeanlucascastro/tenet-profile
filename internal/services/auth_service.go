package service

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) ValidateToken(token string) (bool, error) {

	//TODO: implement call to check token validity with auth service.
	return true, nil
}
