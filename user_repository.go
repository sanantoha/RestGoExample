package main

import (
    "log"
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

type UserRepository interface {
    GetUser(name string) (User, error)
    GetUsers() ([]User, error)
    InsertUser(user *User) error
    UpdateUser(user *User) (bool, error)
    DeleteUser(name string) (bool, error)
}

type UserRepositoryImpl struct {
    db *sql.DB
}

func (ur UserRepositoryImpl) GetUser(name string) (User, error) {
    sqlStatement := `SELECT name, age, last_updatetime FROM users WHERE name = $1`

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
            return user, fmt.Errorf("Error while execute GetUser query: %s, %s", sqlStatement, err)
    }
}

func (ur UserRepositoryImpl) GetUsers() ([]User, error) {
    sqlStatement := `SELECT name, age, last_updatetime FROM users`  

    rows, err := ur.db.Query(sqlStatement)
    if err != nil {
        return nil, fmt.Errorf("Error while execute GetUsers query: %s, %s", sqlStatement, err)
    }

    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.Name, &user.Age, &user.LastUpdatetime)
        if err != nil {
            return nil, fmt.Errorf("Error while execute GetUsers query: %s and fetching data %s", sqlStatement, err)
        }
        users = append(users, user)     
    }
    err = rows.Err()
    if err != nil {
        return nil, fmt.Errorf("Error while execute GetUsers query: %s and encounter during iteration %s", sqlStatement, err)
    }

    log.Println("GetUsers qeury:", sqlStatement, "return", users)
    return users, nil
}

func (ur UserRepositoryImpl) InsertUser(user *User) error {
    sqlStatement := `INSERT INTO users (name, age, last_updatetime) VALUES ($1, $2, $3)`

    _, err := ur.db.Exec(sqlStatement, user.Name, user.Age, user.LastUpdatetime)
    if err != nil {
        return fmt.Errorf("Error while execute InsertUser query %s, error %s", sqlStatement, err)
    }
    log.Println("InsertUser query:", sqlStatement, "user:", *user)
    return nil
}

func (ur UserRepositoryImpl) UpdateUser(user *User) (bool, error) {
    sqlStatement := `UPDATE users SET age = $2, last_updatetime = $3 WHERE name = $1`

    result, err := ur.db.Exec(sqlStatement, user.Name, user.Age, user.LastUpdatetime)
    if err != nil {
        return false, fmt.Errorf("Error while execute UpdateUser query %s, error %s", sqlStatement, err)
    }
    rows, err := result.RowsAffected()
    if err != nil {
        return false, fmt.Errorf("Error while execute DeleteUser query %s, RowsAffected, error %s", sqlStatement, err)
    }
    log.Println("UpdateUser query:", sqlStatement, "user:", *user,"updated rows:", rows)
    return rows > 0, nil
}

func (ur UserRepositoryImpl) DeleteUser(name string) (bool, error) {
    sqlStatement := `DELETE FROM users WHERE name = $1`

    result, err := ur.db.Exec(sqlStatement, name)
    if err != nil {
        return false, fmt.Errorf("Error while execute DeleteUser query %s, error %s", sqlStatement, err)
    }
    rows, err := result.RowsAffected()
    if err != nil {
        return false, fmt.Errorf("Error while execute DeleteUser query %s, RowsAffected, error %s", sqlStatement, err)
    }
    log.Println("DeleteUser query:", sqlStatement, "name", name, "deleted rows:", rows)
    return rows > 0, nil
}