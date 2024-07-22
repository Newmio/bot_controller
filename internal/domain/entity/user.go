package entity

type User struct {
	Id      int    `bson:"id_user"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	UserName  string `bson:"username"`
	Role      string `bson:"role"`
}
