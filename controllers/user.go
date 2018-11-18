package controllers

import (
	"github.com/yinrenxin/joeblog/models"
	"github.com/yinrenxin/joeblog/syserror"
	"github.com/astaxie/beego/logs"
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