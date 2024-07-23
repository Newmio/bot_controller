package mongo

import (
	"bot/internal/domain/entity"
	"bot/pkg"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *mongoRepo) GetBots(limit, offset int64) ([]entity.Bot, error) {
	var bots []entity.Bot
	c := db.db.Collection("bots")

	opt := options.Find()
	opt.SetLimit(limit)
	opt.SetSkip(offset)

	cursor, err := c.Find(context.Background(), bson.M{}, opt)
	if err != nil {
		return nil, pkg.Trace(err)
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &bots); err != nil {
		return nil, pkg.Trace(err)
	}

	return bots, nil
}

func (db *mongoRepo) CreateBot(bot entity.Bot) error {
	var result string
	c := db.db.Collection("bots")

	if err := c.FindOne(context.Background(), bson.M{"login": bot.Login}).Decode(&result); err != nil {
		if err != mongo.ErrNoDocuments {
			return pkg.Trace(err)
		}
	}

	if result != "" {
		return nil
	}

	if _, err := c.InsertOne(context.Background(), bot); err != nil {
		return pkg.Trace(err)
	}

	return nil
}
