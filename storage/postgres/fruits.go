package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/auth/modles"
	"github.com/jasurxaydarov/auth/storage/repoi"
)

type FruitsRepo struct{
	db *pgx.Conn
}

func NewFruitRepo(db *pgx.Conn)repoi.FruitsRepoi{

	return &FruitsRepo{
		db: db,
	}
}

func (f *FruitsRepo)GetFruits(ctx context.Context,req modles.GetList )(*[]modles.Fruits,error){

	var fruit modles.Fruits
	var	fruits []modles.Fruits

	Limit:=req.Limit
	offset:=(req.Page-1)* req.Limit
	
	query:=
	`SELECT 
		*
	 FROM
	 	fruits
	 LIMIT $1 OFFSET $2
	

	`
	row,err:=f.db.Query(ctx,query,Limit,offset)

	if err != nil{
		log.Println("eror on GetFruits",err)
		return nil,err
	}

	for row.Next(){
		row.Scan(
			&fruit.FruitId,
			&fruit.FruitName,
		)

		fruits=append(fruits,fruit)
	}
	return &fruits,nil
}
