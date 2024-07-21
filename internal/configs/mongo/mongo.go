package mongo

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenDb() (*mongo.Client, error) {
	v := viper.New()
	v.AddConfigPath("internal/configs/mongo")
	v.SetConfigName("config")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s",
		v.GetString("login"), v.GetString("password"), v.GetString("host"), v.GetString("port"))))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return client, nil
}
