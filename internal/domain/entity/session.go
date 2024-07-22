package entity

type Session struct {
	User    User
	Command string
}

const(
	BadRequset = "400"
)