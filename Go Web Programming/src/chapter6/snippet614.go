// 使用 Go 语言实现一对多以及多对一关系
package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type MapPost struct {
	Id       int
	Content  string
	Author   string
	Comments []MapComment
}

type MapComment struct {
	Id      int
	Content string
	Author  string
	Post    *MapPost
}

var MapDb *sql.DB

func init() {
	var err error
	MapDb, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp.123 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (comment *MapComment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	MapDb.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func GetPostAndComments(id int) (post MapPost, err error) {
	post = MapPost{}
	post.Comments = []MapComment{}
	err = MapDb.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := MapDb.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := MapComment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

func (post *MapPost) Create() (err error) {
	err = MapDb.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
	return
}

func S614() {
	post := MapPost{Content: "Hello World!", Author: "Sau Sheong"}
	post.Create()

	comment := MapComment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.Create()

	readPost, _ := GetPostAndComments(post.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
