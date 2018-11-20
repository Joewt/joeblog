package controllers

import (
	"github.com/yinrenxin/joeblog/models"
	"github.com/yinrenxin/joeblog/syserror"
	"github.com/astaxie/beego/logs"
	"strings"
	"errors"
	"github.com/yinrenxin/joeblog/common"
)

type UserController struct {
	BaseController
}


// @router /login [post]
func (this *UserController) Login() {
	email := this.GetMushString("email", "邮箱不能为空")
	pwd   := this.GetMushString("password", "密码不能为空")
	id, err := models.QueryUserByEmailAndPwd(email, pwd)

	if err != nil {
		this.Abort500(syserror.New("登录失败",err))
	}

	user, _ := models.QueryUserById(id)
	logs.Info("用户信息",user)
	if err != nil {
		logs.Warn(err)
	}

	this.SetSession(SESSION_USER_KEY,user)

	this.JsonOK("登录成功","/")

	logs.Info("登录成功 用户id：", user)


}

// @router /logout [get]
func (this *UserController) Logout() {
	this.MustLogin()
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)
}

// @router /reg [post]
func (this *UserController) Reg() {
	username := this.GetMushString("username", "用户名不能为空")
	email := this.GetMushString("email","邮箱不能为空")
	pwd   := this.GetMushString("password", "密码不能为空")
	pwd1  := this.GetMushString("password1", "确认密码不能为空")

	if strings.Compare(pwd, pwd1) != 0 {
		this.Abort500(errors.New("两次密码不同"))
	}
	//判断是否有同一个用户
	if models.FindUserByEmail(email) == false{
		this.Abort500(errors.New("已经有该用户了"))
	}

	if common.CheckEmail(email) == false {
		this.Abort500(errors.New("邮箱格式错误"))
	}


	//保存用户信息

	id, err := models.SaveUser(username, email, pwd)
	if err != nil {
		this.Abort500(errors.New("注册失败"))
		logs.Warn("用户注册失败: ", err)
	}

	user, _ := models.QueryUserById(id)

	this.SetSession(SESSION_USER_KEY,user)

	this.JsonOK("注册成功", "/")
}