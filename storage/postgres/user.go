package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/auth/modles"
	"github.com/jasurxaydarov/auth/storage/repoi"
)

type UserRepo struct{
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn)repoi.UserRepoI{

	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo)CreateUser(ctx context.Context,req modles.Account )(error){
	return nil
}