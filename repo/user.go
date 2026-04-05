package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/domain"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/user"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
			name,
			email,
			password
		)
		VALUES (
			:name,
			:email,
			:password
		)
		RETURNING id, created_at
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.CreatedAt); err != nil {
			fmt.Println("Failed to scan user data:", err)
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Failed to retrieve inserted user")
	}

	return &user, nil
}

func (r *userRepo) Login(email, password string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) Find(email string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) GetByID(id string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) GetByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) Update(user domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) Delete(id string) error {
	return nil
}

func (r *userRepo) List() ([]domain.User, error) {
	return nil, nil
}

func (r *userRepo) Status(id string) (bool, error) {
	return false, nil
}
