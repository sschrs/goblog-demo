package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title,Slug string
}

func (category Category) Migrate() {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&category)
}

func (category Category) Add()  {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&category)
}

func (category Category) Get(where ...interface{}) Category {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return category
	}
	db.First(&category,where...)
	return category
}

func (category Category) GetAll(where ...interface{}) []Category {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var categories []Category
	db.Find(&categories,where...)
	return categories
}

func (category Category) Update(column string,value interface{})  {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&category).Update(column,value)
}

func (category Category) Updates(data Category){
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&category).Updates(data)
}

func (category Category) Delete()  {
	db,err := gorm.Open(mysql.Open(Dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&category,category.ID)
}
