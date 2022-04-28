package model

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) error {
	stmt := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

	fmt.Println("!!!! in create user", user.Name, user.Email, user.Password)

	_, err := db.Exec(stmt, user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Println("я ОШИБКА!!!!!!!!!!!", err)
	}
	fmt.Println("я не ОШИБКА!!!!!!!!!!!")
	return nil
}

func GetUser(userID string) (*User, error) {
	var user *User

	stmt := `SELECT id, name, email, password FROM users WHERE id = ?`

	if err := db.QueryRow(stmt, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("users not exists with id: " + userID)
		}

		return nil, err
	}

	return user, nil
}
