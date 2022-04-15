package sqlstore_test

import (
	"testing"

	"github.com/FreezeOMatic/Firsttry/internal/app/models"
	"github.com/FreezeOMatic/Firsttry/internal/app/store"
	"github.com/FreezeOMatic/Firsttry/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := models.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := models.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}