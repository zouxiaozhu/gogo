package models

import "time"

type Book struct {
	Id int 	`db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Author string `db:"author" json:"author"`
	Price string `db:"price" json:"price"`
	Create_time time.Time `db:"create_time" json:"-"`
	Update_time time.Time `db:"update_time" json:"-"`
}


type BookAdd struct {
	Name string `db:"name" form:"name" binding:"required"`
	Author string `db:"author" form:"author" binding:"required"`
	Price string `db:"price" form:"price" binding:"required"`
}