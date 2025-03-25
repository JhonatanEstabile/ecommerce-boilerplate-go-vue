package user

import (
	"errors"
	"log"
)

type Service struct {
	repo        Repository
	authService AuthAdapter
}

func NewService(r Repository, a AuthAdapter) *Service {
	return &Service{
		repo:        r,
		authService: a,
	}
}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("email ou senha inválidos")
	}

	if !s.authService.ComparePassword(user.Password, password) {
		return "", errors.New("email ou senha inválidos")
	}

	token, err := s.authService.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("erro ao gerar token")
	}

	return token, nil
}

func (s *Service) Create(user User) error {
	hashedPassword, err := s.authService.HashPassword(user.Password)
	if err != nil {
		return errors.New("Invalid password")
	}

	user.Password = hashedPassword

	err = s.repo.Create(user)
	if err != nil {
		log.Println("Error to save user in database", err.Error())
		return errors.New("Erro to save the new user in database")
	}

	return nil
}
