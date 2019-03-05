package models

import (
	"appleforum/models/db"
	"database/sql"
	"errors"
)

// User model
type User struct {
	Name   string
	Passwd string
	Email  string
}

// CreateUser create user
func CreateUser(name, passwd, email string) error {
	// check if user exists
	var selectName string
	err := db.QueryRow("SELECT name FROM users WHERE name = $1", name).Scan(&selectName)
	if err != nil {
		if err == sql.ErrNoRows {
			db.Exec("INSERT INTO users (name, passwd, email) VALUES ($1, $2, $3)",
				name, passwd, email)
			return nil
		}
		return errors.New("user already exists")
	}
	return errors.New("create user error")
}
