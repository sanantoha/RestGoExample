package main

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	GetUser(name string) (User, error)
	// GetUsers() ([]User, error)
	// InsertUser(user User) error
	// UpdateUser(user User) error
	// DeleteUser(user User) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func (ur UserRepositoryImpl) GetUser(name string) (User, error) {
	sqlStatement := "SELECT name, age, last_updatetime FROM users WHERE name = $1"

	var user User

	row := ur.db.QueryRow(sqlStatement, name)
	err := row.Scan(&user.Name, &user.Age, &user.LastUpdatetime)

	log.Println("qeury:", sqlStatement, "with params:", name, "return", user)

	switch err {
		case sql.ErrNoRows:
			log.Println("No rows were returned!")
			return user, fmt.Errorf("No rows were returned!")
		case nil:
			return user, nil
		default:
			return user, fmt.Errorf("Error while execute GetUser query: %s", err)
	}
}