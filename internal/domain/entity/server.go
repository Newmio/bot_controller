package entity

type BotServer struct {
	Name     string `bson:"name"`
	Host     string `bson:"host"`
	IP       string `bson:"ip"`
	BotCount int    `bson:"bot_count"`
}
