package user

import "github.com/mhbhuiyan99/Finance-Dashboard-System/domain"

type Service interface {

}

type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	GetByID(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	List() ([]domain.User, error)
	Delete(id string) error
	Update(u domain.User) (*domain.User, error)
	Status(id string) (bool, error)
}