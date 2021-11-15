// Gorm
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type GormPost struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []GormComment
	CreatedAt time.Time
}

type GormComment struct {
	Id         int
	Content    string
	Author     string `sql:"not null"`
	GormPostId int    `sql:"index"`
	CreatedAt  time.Time
}

var GormDb *gorm.DB

func init() {
	var err error
	GormDb, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp.123 sslmode=disable")
	if err != nil {
		panic(err)
	}
	GormDb.DropTableIfExists(&GormPost{}, &GormComment{})
	GormDb.AutoMigrate(&GormPost{}, &GormComment{})
}

func main() {
	post := GormPost{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)

	GormDb.Create(&post)
	fmt.Println(post)

	comment := GormComment{Content: "Good post!", Author: "Joe"}
	GormDb.Model(&post).Association("Comments").Append(comment)

	var readPost GormPost
	GormDb.Where("author = $1", "Sau Sheong").First(&readPost)
	fmt.Println(readPost)
	var comments []GormComment
	GormDb.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
