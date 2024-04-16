package storage

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/Corray333/dating/internal/domains/date/types"
	user_types "github.com/Corray333/dating/internal/domains/user/types"
	"github.com/Corray333/dating/internal/storage"
	sq "github.com/Masterminds/squirrel"
	"github.com/gorilla/websocket"
)

type DateStorage struct {
	storage.Storage
	Dates          map[int]types.Date
	UsersSearching map[int]*websocket.Conn
}

func (s *DateStorage) StartSearching(id int, conn *websocket.Conn) error {

	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}

	// Close connection if searching from other device
	if v, ok := s.UsersSearching[id]; ok {
		v.Close()
	}
	var user user_types.User
	rows, err := tx.Queryx("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return fmt.Errorf("user not found")
	}
	if err := rows.StructScan(&user); err != nil {
		return err
	}
	// Try to find chat partner in searching_users
	query := sq.Select("*").From("searching_users")
	sex := []int{}
	if user.Search&user_types.SearchFemale == 1 {
		sex = append(sex, user_types.SexFemale)
	}
	if user.Search&user_types.SearchMale == 1 {
		sex = append(sex, user_types.SexMale)
	}
	query = query.Where(sq.Eq{"sex": sex})

	rows, err = tx.Queryx(query.ToSql())
	if err != nil {
		return err
	}
	for rows.Next() {
		var partner user_types.User
		if err := rows.StructScan(&partner); err != nil {
			slog.Error(err.Error())
			continue
		}

	}

	if _, err := tx.Queryx("INSERT INTO searching_users VALUES ($1, $2, $3, $4, $5, $6)",
		id, user.City, user.Sex, user.Orientation, user.Birth, user.Search); err != nil {
		return err
	}

	// Add user to the map
	mu := sync.Mutex{}
	mu.Lock()
	s.UsersSearching[id] = conn
	mu.Unlock()

	tx.Commit()

	return nil
}
