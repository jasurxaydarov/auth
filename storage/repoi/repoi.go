package repoi

import (
	"context"

	"github.com/jasurxaydarov/auth/modles"
)

type FruitsRepoi interface{

	GetFruits(ctx context.Context,req modles.GetList )(*[]modles.Fruits,error)
}
type UserRepoI interface{

	CreateUser(ctx context.Context,req modles.Account )(error)
}