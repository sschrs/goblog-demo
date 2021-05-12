package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title,Slug,Description,Content,Picture_url string
	CategoryID int
}

func (post Post) Migrate() {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&post)
}

func (post Post) Add()  {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&post)
}

func (post Post) Get(where ...interface{}) Post {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return post
	}
	db.First(&post,where...)
	return post
}

func (post Post) GetAll(where ...interface{}) []Post {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var posts []Post
	db.Find(&posts,where...)
	return posts
}

func (post Post) Update(column string,value interface{})  {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&post).Update(column,value)
}

func (post Post) Updates(data Post){
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&post).Updates(data)
}

func (post Post) Delete()  {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&post,post.ID)
}
