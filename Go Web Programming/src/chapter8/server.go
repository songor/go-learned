package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"path"
	"strconv"
)

type Text interface {
	fetch(id int) (err error)
}

type TPost struct {
	Db      *sql.DB
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (post *TPost) fetch(id int) (err error) {
	err = post.Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func server() {
	var err error
	db, err := sql.Open("postgres", "user=gwp dbname=gwp password=gwp.123 sslmode=disable")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/post/", handleRequest(&TPost{Db: db}))
	server.ListenAndServe()
}

func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// curl -i -X GET http://127.0.0.1:8080/post/1
func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
