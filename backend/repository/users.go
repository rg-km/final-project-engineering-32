package repository

import (
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {
	//beginanswer
	var user User
	err := u.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}
	return user, nil
	//endanswer return User{}, nil
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	//beginanswer
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
	//endanswer return []User{}, nil
}

func (u *UserRepository) Login(username string, password string) (*string, error) {
	//beginanswer
	var user User
	err := u.db.QueryRow("SELECT username FROM user WHERE username = ? AND password = ?", username, password).Scan(&user.Username)
	if err != nil {
		return nil, errors.New("Login Failed")
	}
	return &user.Username, nil
	//endanswer return nil, nil
}

func (u *UserRepository) InsertUser(username string, password string, role string) error {
	//beginanswer
	_, err := u.db.Exec("INSERT INTO user (username, password, role) VALUES (?, ?, ?)", username, password, role)
	if err != nil {
		return err
	}
	return nil
	//endanswer return nil
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {
	//beginanswer
	var user User
	err := u.db.QueryRow("SELECT role FROM user WHERE username = ?", username).Scan(&user.Role)
	if err != nil {
		return nil, err
	}
	return &user.Role, nil
	//endanswer return nil, nil
}
