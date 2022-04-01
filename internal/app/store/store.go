package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = db.Exec(selectQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	defer db.Close()

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
