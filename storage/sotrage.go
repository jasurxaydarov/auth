package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/auth/storage/postgres"
	"github.com/jasurxaydarov/auth/storage/repoi"
)
type StorageI interface{
	FruitsRepo()repoi.FruitsRepoi
	UserRepo()repoi.UserRepoI
}

type storage  struct{
	userRepo repoi.UserRepoI
	fruitRepo repoi.FruitsRepoi
}


func NewStorage(db *pgx.Conn)StorageI{

	return &storage{
		userRepo: postgres.NewUserRepo(db),
		fruitRepo: postgres.NewFruitRepo(db),
	}
}

func (s *storage)FruitsRepo()repoi.FruitsRepoi{
	return s.fruitRepo
}
func (s *storage)UserRepo()repoi.UserRepoI{
	return s.userRepo
}


