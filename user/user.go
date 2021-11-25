package user

import (
	"time"
)

type User struct{
	ID 			int
	Email		string
	Password	string
	Nama		string
	UserImg		string
	Created		time.Time
	Updated		time.Time 
	Role		string
	Pekerjaan	string
}
