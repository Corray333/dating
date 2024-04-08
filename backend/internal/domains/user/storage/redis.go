package storage

import (
	"strconv"

	"github.com/Corray333/dating/pkg/server/auth"
)

func (s *UserStorage) SetRefreshToken(id int, agent string, refresh string) error {
	if err := s.Redis.Set(ctx, strconv.Itoa(id)+agent, refresh, auth.RefreshTokenLifeTime).Err(); err != nil {
		return err
	}
	return nil
}
