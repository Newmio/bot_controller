package redis

import (
	"bot/pkg"
	"context"
	"fmt"
)

func (db *redisRepo) DeleteCommandActions(userId int, command string) error {
	return db.db.Del(context.Background(), fmt.Sprintf("session_actions:%d-%s", userId, command)).Err()
}

func (db *redisRepo) GetCommandActions(userId int, command string) (map[string]string, error) {
	return db.db.HGetAll(context.Background(), fmt.Sprintf("session_actions:%d-%s", userId, command)).Result()
}

func (db *redisRepo) SetCommandActions(userId int, command string, actions map[string]string) error {
	return db.db.HSet(context.Background(), fmt.Sprintf("session_actions:%d-%s", userId, command), actions).Err()
}

func (db *redisRepo) GetSessionCommand(userId int) (string, error) {

	resp, err := db.db.HGet(context.Background(), fmt.Sprintf("session:%d", userId), "command").Result()
	if err != nil {
		return "", pkg.Trace(err)
	}

	return resp, nil
}

func (db *redisRepo) SetSessionCommand(userId int, command string) error {
	return db.db.HSet(context.Background(), fmt.Sprintf("session:%d", userId), "command", command).Err()
}
