package user

import (
	"github.com/mhbhuiyan99/Finance-Dashboard-System/domain"
)

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (s *service) Create(u domain.User) (*domain.User, error) {
	u.Role = "viewer" // default role for new users
	u.IsActive = true // set new users as active by default
	
	usr, err := s.usrRepo.Create(u)

	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}

func (s *service) Find(email string) (*domain.User, error) {
	usr, err := s.usrRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}

func (s *service) Login(email, password string) (*domain.User, error) {
	return s.usrRepo.Login(email, password)
}

