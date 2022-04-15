package teststore

import (
	"github.com/FreezeOMatic/Firsttry/internal/app/models"
	"github.com/FreezeOMatic/Firsttry/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*models.User
}

func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}