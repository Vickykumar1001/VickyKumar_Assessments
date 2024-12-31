package repository

import (
	"database/sql"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (repo *AuthRepository) CreateUser(username, password string) error {
	_, err := repo.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}
