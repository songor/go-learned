// Sqlx
package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SqlxPost struct {
	Id         int
	Content    string
	AuthorName string `db:"author"`
}

var SqlxDb *sqlx.DB

func init() {
	var err error
	SqlxDb, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=gwp.123 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetSqlxPost(id int) (post SqlxPost, err error) {
	post = SqlxPost{}
	err = SqlxDb.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

func (post *SqlxPost) Create() (err error) {
	SqlxDb.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func S617() {
	post := SqlxPost{Content: "Hello World", AuthorName: "Sau Sheong"}
	post.Create()
	readPost, _ := GetSqlxPost(post.Id)
	fmt.Println(readPost)
}
