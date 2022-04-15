package teststore

import (
	"github.com/FreezeOMatic/Firsttry/internal/app/models"
	"github.com/FreezeOMatic/Firsttry/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*models.User),
	}
	return s.userRepository
}
