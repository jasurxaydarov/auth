package main

import (
	"fmt"
	"log"


	"github.com/jasurxaydarov/auth/config"
	"github.com/jasurxaydarov/auth/pgx/db"
)

func main(){

	cfg:=config.Load()

	pgxConn,err:=db.ConnToDb(cfg.PgConfig)
	if err!= nil{
		log.Println("error on  con to db ",err)
		return
	}
	fmt.Println(pgxConn)
}