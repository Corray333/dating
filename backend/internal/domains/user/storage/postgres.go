package storage

import (
	"context"
	"fmt"

	"github.com/Corray333/dating/internal/domains/user/types"
	"github.com/Corray333/dating/internal/storage"
	"github.com/Corray333/dating/pkg/server/auth"
	_ "github.com/lib/pq"
)

var ctx = context.Background()

type UserStorage storage.Storage

// InsertUser inserts a new user into the database and returns the id
func (s *UserStorage) InsertUser(user types.User, agent string) (int, string, error) {
	passHash, err := auth.Hash(user.Password)
	if err != nil {
		return -1, "", err
	}
	user.Password = passHash

	rows := s.DB.QueryRow(`
		INSERT INTO users (username, email, password, name, surname, patronymic, sex, referal, orientation_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING user_id;
	`, user.Username, user.Email, user.Password, user.Name, user.Surname, user.Patronymic, user.Sex, user.Referal, user.OrientationID)

	if err := rows.Scan(&user.ID); err != nil {
		return -1, "", err
	}

	refresh, err := auth.CreateToken(user.ID, auth.RefreshTokenLifeTime)
	if err != nil {
		return -1, "", err
	}

	// _, err = s.DB.Queryx(`
	// 	INSERT INTO user_token (user_id, token) VALUES ($1, $2);
	// `, user.ID, refresh)
	// if err != nil {
	// 	return -1, "", err
	// }

	if err := s.SetRefreshToken(user.ID, agent, refresh); err != nil {
		return -1, "", err
	}

	return user.ID, refresh, nil
}

// LoginUser checks if the user exists and the password is correct
func (s *UserStorage) LoginUser(user types.User, agent string) (int, string, error) {
	password := user.Password

	rows := s.DB.QueryRow(`
		SELECT user_id, password FROM users WHERE email = $1;
	`, user.Email)

	if err := rows.Scan(&user.ID, &user.Password); err != nil {
		return -1, "", err
	}
	if !auth.Verify(user.Password, password) {
		return -1, "", fmt.Errorf("invalid password")
	}

	// Auto update refresh token
	refresh, err := auth.CreateToken(user.ID, auth.RefreshTokenLifeTime)
	if err != nil {
		return -1, "", err
	}

	if err := s.SetRefreshToken(user.ID, agent, refresh); err != nil {
		return -1, "", err
	}

	return user.ID, refresh, nil
}

// CheckAndUpdateRefresh checks if the refresh token is valid and updates it
func (s *UserStorage) CheckAndUpdateRefresh(id int, refresh string) (string, error) {
	rows, err := s.DB.Queryx(`
		SELECT token FROM user_token WHERE user_id = $1 AND token = $2;
	`, id, refresh)
	if err != nil {
		return "", err
	}
	if !rows.Next() {
		return "", fmt.Errorf("invalid refresh token")
	}
	newRefresh, err := auth.CreateToken(id, auth.RefreshTokenLifeTime)
	if err != nil {
		return "", err
	}
	_, err = s.DB.Queryx(`
		UPDATE user_token SET token = $1 WHERE user_id = $2;
	`, newRefresh, id)
	if err != nil {
		return "", err
	}
	return newRefresh, nil
}

func (s *UserStorage) SelectUser(id string) (types.User, error) {
	var user types.User
	rows, err := s.DB.Queryx(`
		SELECT * FROM users WHERE user_id = $1;
	`, id)
	if err != nil {
		return user, err
	}
	if !rows.Next() {
		return user, fmt.Errorf("user not found")
	}
	if err := rows.StructScan(&user); err != nil {
		return user, err
	}
	user.Password = ""
	return user, nil
}

func (s *UserStorage) UpdateUser(user types.User) error {
	fmt.Println()
	fmt.Println(user)
	fmt.Println()
	_, err := s.DB.Queryx(`
		UPDATE users SET username = $1, email = $2, avatar = $3 WHERE user_id = $4;
	`, user.Username, user.Email, user.Avatar, user.ID)
	return err
}
