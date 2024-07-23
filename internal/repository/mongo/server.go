package mongo

import (
	"bot/internal/domain/entity"
	"bot/pkg"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *mongoRepo) CreateServer(server entity.BotServer) error {
	var check entity.BotServer
	c := db.db.Collection("bot_servers")

	if err := c.FindOne(context.Background(), bson.M{"ip": server.IP}).Decode(&check); err != nil {
		if err != mongo.ErrNoDocuments {
			return pkg.Trace(err)
		}
	}

	if check.IP != "" {
		return nil
	}

	if _, err := c.InsertOne(context.Background(), server); err != nil {
		return pkg.Trace(err)
	}

	return nil
}

func (db *mongoRepo) GetServers(limit, offset int64) ([]entity.BotServer, error) {
	var servers []entity.BotServer
	c := db.db.Collection("bot_servers")

	opt := options.Find()
	opt.SetLimit(limit)
	opt.SetSkip(offset)

	cursor, err := c.Find(context.Background(), bson.M{}, opt)
	if err != nil {
		return nil, pkg.Trace(err)
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &servers); err != nil {
		return nil, pkg.Trace(err)
	}

	return servers, nil
}
