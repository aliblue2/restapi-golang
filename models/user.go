package models

import "azno-space.com/azno/db"

type User struct {
	Id       int64
	Name     string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func SignupUser(email, name, password string) (int64, error) {
	query := `INSERT INTO users (email,name,password) VALUES (?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(email, name, password)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (u *User) ValidateCreadintials(email string) (string, error) {
	query := `SELECT id , password FROM users WHERE email = ?`

	var retirivedPassword string
	sqlRow := db.DB.QueryRow(query, email)

	err := sqlRow.Scan(&u.Id, &retirivedPassword)

	return retirivedPassword, err

}
