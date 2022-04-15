package store

import "github.com/FreezeOMatic/Firsttry/internal/app/models"

type UserRepository interface {
	Create(*models.User) error
	FindByEmail(string) (*models.User, error)
}
