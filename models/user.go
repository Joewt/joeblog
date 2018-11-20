package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/yinrenxin/joeblog/common"
	"github.com/yinrenxin/joeblog/syserror"
	"github.com/astaxie/beego/logs"
)

var DB orm.Ormer

func init() {
	DB = orm.NewOrm()
	DB.Using("default") // 默认使用 default，你可以指定为其他数据库
	var maps []orm.Params
	num, _ := DB.Raw("SELECT * FROM user").Values(&maps)
	if num == 0 {
		user := new(User)
		user.Name = "joe"
		user.Email = "joewt@foxmail.com"
		user.Pwd = common.MD5("123456")
		user.Avatar = "/static/images/item.png"
		user.Role = 0
		DB.Insert(user)
		logs.Info("管理员信息插入成功")
	} else {
		logs.Info("管理员信息已经添加")
	}
}


func QueryUserByEmailAndPwd(email, pwd string) (uint,error) {
	user := User{Email: email}
	err := DB.Read(&user,"Email")
	if err != nil{
		err = syserror.New("没有该账号",nil)
		return 247, err
	}
	if user.Pwd != common.MD5(pwd) {
		err = syserror.New("密码错误", nil)
		return 247, err
	}
	return user.Id, nil
}

func QueryUserById(id uint) (user User, err error){
	user = User{Id: id}
	err = DB.Read(&user, "Id")
	if err != nil {
		logs.Warn(err)
	}
	return user, nil
}


func FindUserByEmail(email string) bool {
	user := User{Email: email}
	err := DB.Read(&user, "Email")
	if err != nil {
		return true
	}
	return false
}


func SaveUser(username,email, pwd string) (uint, error) {
	user := new(User)
	user.Name = username
	user.Pwd = common.MD5(pwd)
	user.Email = email
	user.Avatar = "/static/images/item.png"
	user.Role = 1
	id, err := DB.Insert(user)
	if err != nil {
		return uint(id), err
	}
	logs.Info("新注册一个用户: ", id)
	return uint(id), nil
}