package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username,Password string
}

func (user User) Migrate() {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&user)
}

func (user User) Add()  {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&user)
}

func (user User) Get(where ...interface{}) User {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return user
	}
	db.First(&user,where...)
	return user
}

func (user User) GetAll(where ...interface{}) []User {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var users []User
	db.Find(&users,where...)
	return users
}

func (user User) Update(column string,value interface{})  {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&user).Update(column,value)
}

func (user User) Updates(data User){
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&user).Updates(data)
}

func (user User) Delete()  {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&user,user.ID)
}
