package redis

import (
	"bot/internal/domain/entity"
	"bot/pkg"
	"context"
	"encoding/json"
	"fmt"
)


func (db *redisRepo) GetSession(userId int) (entity.Session, error) {
	var session entity.Session

	resp, err := db.db.HGetAll(context.Background(), fmt.Sprintf("session:%d", userId)).Result()
	if err != nil {
		return session, pkg.Trace(err)
	}

	if len(resp) == 0 {
		return session, nil
	}

	if err := json.Unmarshal([]byte(resp["data"]), &session); err != nil {
		return session, pkg.Trace(err)
	}

	return session, nil
}

func (db *redisRepo) CreateSession(session entity.Session) error {
	model, err := json.Marshal(session)
	if err != nil {
		return pkg.Trace(err)
	}

	if err := db.db.HSet(context.Background(), fmt.Sprintf("session:%d", session.User.Id), "data", string(model)).Err(); err != nil {
		return pkg.Trace(err)
	}

	return nil
}
