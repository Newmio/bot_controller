package entity

type Bot struct {
	UserId int    `bson:"id_user"`
	Login  string `bson:"login"`
	Pass   string `bson:"pass"`
}
