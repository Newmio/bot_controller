package mongo

import (
	"bot/internal/domain/entity"
	"bot/pkg"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *mongoRepo) CreateUser(user entity.User) error {
	var check entity.User
	c := db.db.Collection("users")

	if err := c.FindOne(context.Background(), bson.M{"tg_id": user.TgId}).Decode(&check); err != nil {
		if err != mongo.ErrNoDocuments {
			return pkg.Trace(err)
		}
	}

	if check.TgId != 0 {
		return nil
	}

	if _, err := c.InsertOne(context.Background(), user); err != nil{
		return pkg.Trace(err)
	}

	return nil
}
