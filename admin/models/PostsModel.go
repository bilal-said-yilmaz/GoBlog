package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



type Post struct {
	gorm.Model
	Title, Slug, Description, Content, PictureURL string
	CategoryID                                    int
}

func (post Post) Migrate() {
	Db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		return
	}
	Db.AutoMigrate(&post)
}

func (post Post) Add() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		return
	}
	db.Create(&post)
}

func (post Post) Get(where ...interface{}) Post {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		return post
	}
	db.First(&post, where...)
	return post
}

func (post Post) GetAll(where ...interface{}) []Post {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		return nil
	}
	var posts []Post
	db.Find(&posts, where...)

	return posts
}

func (post Post) Update(column string, value interface{}) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		return
	}
	db.Model(&post).Update(column, value)
}

func (post Post) Updates(data Post) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		return
	}
	db.Model(&post).Updates(data)
}

func (post Post) Delete() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		return
	}
	db.Delete(&post, post.ID)
}
