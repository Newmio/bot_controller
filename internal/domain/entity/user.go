package entity

type User struct {
	TgId      int    `bson:"tg_id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	UserName  string `bson:"username"`
	Role      string `bson:"role"`
}
