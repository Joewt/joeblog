package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type User struct {
	Id     uint
	Name   string `orm:"unique"`
	Email  string `orm:"unique"`
	Pwd    string
	Avatar string
	Role   int    `orm:"default(1)"` //0管理员 1普通用户
}

type Article struct {
	Id uint
	Key string
	User_id uint
	Title string
	Summary string
	Content string
	Praise int     `orm:"default(0)"`
	Discuss int    `orm:"default(0)"`
	CreatedAt time.Time `orm:"auto_now_add;type(date)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(date)"`
	DeletedAt time.Time `orm:"auto_now_add;type(date)"`
}

type Comment struct {
	Id uint
	Article_id uint
	User_id uint
	Content string
	CreatedAt time.Time
}


func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User),new(Article),new(Comment))
	orm.RunSyncdb("default", false, true)
}
