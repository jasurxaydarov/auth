package modles

import (
	"github.com/google/uuid"
)

type UserReg struct {
	UserId   uuid.UUID `json:"user_id"`
	UserName string    `json:"user_name"`
	Password string    `json:"password"`
	Gmail    string    `json:"gamil"`
}

type CheckOtp struct {
	Code  string `json:"code"`
	Gmail string `json:"gmail"`
}

type Account struct {
	UserId   uuid.UUID `json:"user_id"`
	UserName string    `json:"user_name"`
	Password string    `json:"password"`
	Gmail    string    `json:"gamil"`
}

type RespLogIn struct {
	Token string 		`json:"token"`
}

type GetList struct{
	Limit int
	Page int
}

type Fruits struct{
	FruitId	  string	`json:"fruit_id"`
	FruitName string 	`json:"fruit_name"`

}