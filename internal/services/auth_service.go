package service

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) ValidateToken(token string) (bool, error) {

	return true, nil
}
