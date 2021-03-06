// 在内存里面存储数据
package main

import "fmt"

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func S61() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(post)
	post = Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	store(post)
	post = Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	store(post)
	post = Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}
	store(post)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
