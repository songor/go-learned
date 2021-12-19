package main

import (
	"database/sql"
	_ "github.com/jmoiron/sqlx"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp.123 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

//func retrieve(id int) (post Post, err error) {
//
//}
