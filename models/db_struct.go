package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type Base struct {
	Id uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type User struct {
	Id uint
	Name        string  `orm:"unique"`
	Email		string	`orm:"unique"`
	Pwd			string
	Avatar		string
	Role        int `orm:"default(1)"`  //0管理员 1普通用户
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Article struct {

}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
