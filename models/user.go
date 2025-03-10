package models

import "github.com/taylorjo02/go-RESTAPI/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err 
	}

	defer statement.Close()

	result, err := statement.Exec(u.Email, u.Password)
	if err != nil {
		return err 
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return err 
	}

	u.ID = userId

  return err
}
