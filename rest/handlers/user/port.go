package user

import "github.com/mhbhuiyan99/Finance-Dashboard-System/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}